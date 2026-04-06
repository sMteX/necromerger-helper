import React, { useState, useEffect } from 'react';
import type { Plan, CalculationResponse, ExperimentID, LegendaryID, RuneType } from './types/game';
import { ExperimentCard } from './components/ExperimentCard';
import { LegendaryCard } from './components/LegendaryCard';
import { experimentDescriptions, experimentMaxLevels, legendaries, runeTypes, devourerLevels } from './data';

const createEmptyCounts = <T extends string>(keys: T[]): Record<T, number> => {
  return keys.reduce((acc, key) => ({ ...acc, [key]: 0 }), {} as Record<T, number>);
};

const initialPlan: Plan = {
  id: 1,
  name: "Current Run",
  devourerLevel: 35,
  featTiers: 0,
  otherMultiplier: 0,
  groupBonusCount: 1,
  leftoverShards: 0,
  legendaryCounts: createEmptyCounts(legendaries.map(l => l.id)),
  experimentLevels: createEmptyCounts(Object.keys(experimentDescriptions) as ExperimentID[]),
  possessedRunes: createEmptyCounts(runeTypes),
  possessedLegendaries: createEmptyCounts(legendaries.map(l => l.id)),
  notes: ""
};

const App: React.FC = () => {
  const [plan, setPlan] = useState<Plan>(initialPlan);
  const [results, setResults] = useState<CalculationResponse | null>(null);
  const [activeTab, setActiveTab] = useState<'experiments' | 'legendaries' | 'resources'>('experiments');

  useEffect(() => {
    const timer = setTimeout(async () => {
      try {
        const resp = await fetch('/api/recalculate', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(plan)
        });
        if (resp.ok) {
          const data = await resp.json();
          setResults(data);
        }
      } catch (err) {
        console.error("Failed to recalculate", err);
      }
    }, 300);

    return () => clearTimeout(timer);
  }, [plan]);

  const updatePlan = (patch: Partial<Plan>) => {
    setPlan(prev => ({ ...prev, ...patch }));
  };

  const updateExpLevel = (id: ExperimentID, level: number) => {
    updatePlan({ experimentLevels: { ...plan.experimentLevels, [id]: level } });
  };

  const updateLegendaryCount = (id: LegendaryID, count: number) => {
    updatePlan({ legendaryCounts: { ...plan.legendaryCounts, [id]: count } });
  };

  const updatePossessedRune = (rune: RuneType, count: number) => {
    updatePlan({ possessedRunes: { ...plan.possessedRunes, [rune]: count } });
  };

  const updatePossessedLegendary = (id: LegendaryID, count: number) => {
    updatePlan({ possessedLegendaries: { ...plan.possessedLegendaries, [id]: count } });
  };

  const pre100Exps: ExperimentID[] = [
    'seasoning', 'strength', 'taste', 'capacity', 'body_snatcher', 
    'weakening', 'damage_cap', 'ice_chest', 'poison_chest', 
    'blood_chest', 'moon_chest', 'death_chest', 'cosmic_chest'
  ];

  const post100Exps: ExperimentID[] = [
    'seasoning_2', 'strength_2', 'taste_2', 'capacity_2'
  ];

  const handleReset = () => {
    if (confirm("Are you sure you want to reset the entire plan?")) {
      setPlan(initialPlan);
    }
  };

  const handleSave = () => {
    localStorage.setItem('necro_prestige_plan', JSON.stringify(plan));
    alert("Plan saved to local storage!");
  };

  const handleLoad = () => {
    const saved = localStorage.getItem('necro_prestige_plan');
    if (saved) {
      setPlan(JSON.parse(saved));
    } else {
      alert("No saved plan found.");
    }
  };

  return (
    <div className="min-h-screen bg-slate-900 text-white flex flex-col font-sans selection:bg-indigo-500/30">
      <header className="sticky top-0 z-50 w-full border-b border-slate-800 bg-slate-950/80 backdrop-blur-xl">
        <div className="max-w-[1600px] mx-auto flex h-20 items-center justify-between px-6">
          <div className="flex items-center gap-4">
            <div className="bg-indigo-600/20 p-2.5 rounded-2xl border border-indigo-500/20 shadow-lg shadow-indigo-500/5">
              <img src="/assets/images/time_shard.png" alt="Logo" className="w-8 h-8"/>
            </div>
            <div>
              <h1 className="text-xl font-black tracking-tight uppercase leading-none mb-1">Necro Prestige</h1>
              <div className="text-[10px] font-black text-slate-500 uppercase tracking-widest flex items-center gap-2">
                <span className="w-1.5 h-1.5 rounded-full bg-emerald-500 animate-pulse"></span>
                Simulation Engine
              </div>
            </div>
          </div>
          
          <div className="flex items-center gap-3">
            <button onClick={handleReset} className="px-4 py-2 bg-slate-900 hover:bg-red-500/10 hover:text-red-400 border border-slate-800 rounded-xl text-xs font-black uppercase tracking-widest transition-all">Reset</button>
            <button onClick={handleLoad} className="px-4 py-2 bg-slate-900 hover:bg-slate-800 border border-slate-800 rounded-xl text-xs font-black uppercase tracking-widest transition-all">Load</button>
            <button onClick={handleSave} className="px-6 py-2 bg-indigo-600 hover:bg-indigo-500 border border-indigo-400/20 rounded-xl text-xs font-black uppercase tracking-widest transition-all shadow-lg shadow-indigo-600/20">Save Plan</button>
          </div>
        </div>
      </header>

      <div className="flex-1 flex overflow-hidden max-w-[1600px] mx-auto w-full">
        <main className="flex-1 overflow-y-auto p-8 custom-scrollbar">
          <div className="space-y-12">
            {/* General Inputs Section */}
            <section className="bg-slate-800/20 border border-slate-700/30 rounded-[32px] p-8">
              <h2 className="text-xs font-black text-slate-500 uppercase tracking-[0.2em] mb-8 flex items-center gap-3">
                <span className="w-8 h-px bg-slate-700"></span> General Parameters
              </h2>
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
                <div className="space-y-3">
                  <label className="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Devourer Level</label>
                  <select 
                    value={plan.devourerLevel} 
                    onChange={(e) => updatePlan({ devourerLevel: parseInt(e.target.value) })}
                    className="w-full bg-slate-950 border border-slate-700 rounded-2xl px-4 py-3 text-sm font-bold text-white focus:ring-2 focus:ring-indigo-500 outline-none appearance-none cursor-pointer"
                  >
                    {devourerLevels.map(lvl => <option key={lvl} value={lvl}>Level {lvl}</option>)}
                  </select>
                </div>
                <div className="space-y-3">
                  <label className="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Feat Tiers (+10% each)</label>
                  <input 
                    type="number" 
                    value={plan.featTiers} 
                    onChange={(e) => updatePlan({ featTiers: parseInt(e.target.value) || 0 })}
                    className="w-full bg-slate-950 border border-slate-700 rounded-2xl px-4 py-3 text-sm font-bold text-white focus:ring-2 focus:ring-indigo-500 outline-none"
                  />
                </div>
                <div className="space-y-3">
                  <label className="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Group Bonuses (Limit)</label>
                  <input 
                    type="number" 
                    value={plan.groupBonusCount} 
                    onChange={(e) => updatePlan({ groupBonusCount: parseInt(e.target.value) || 0 })}
                    className="w-full bg-slate-950 border border-slate-700 rounded-2xl px-4 py-3 text-sm font-bold text-white focus:ring-2 focus:ring-indigo-500 outline-none"
                  />
                </div>
                <div className="space-y-3">
                  <label className="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Other Multiplier (%)</label>
                  <input 
                    type="number" 
                    step="0.01"
                    value={plan.otherMultiplier * 100} 
                    onChange={(e) => updatePlan({ otherMultiplier: (parseFloat(e.target.value) || 0) / 100 })}
                    className="w-full bg-slate-950 border border-slate-700 rounded-2xl px-4 py-3 text-sm font-bold text-white focus:ring-2 focus:ring-indigo-500 outline-none"
                  />
                </div>
              </div>
            </section>

            {/* Navigation Tabs */}
            <div className="flex gap-2 p-1 bg-slate-950 border border-slate-800 rounded-2xl w-fit">
              {(['experiments', 'legendaries', 'resources'] as const).map(tab => (
                <button
                  key={tab}
                  onClick={() => setActiveTab(tab)}
                  className={`px-6 py-2.5 rounded-xl text-xs font-black uppercase tracking-widest transition-all ${activeTab === tab ? 'bg-indigo-600 text-white shadow-lg shadow-indigo-600/20' : 'text-slate-500 hover:text-slate-300 hover:bg-slate-900'}`}
                >
                  {tab}
                </button>
              ))}
            </div>

            {/* Experiments Tab */}
            {activeTab === 'experiments' && (
              <div className="space-y-16 animate-in fade-in slide-in-from-bottom-4 duration-500">
                <section>
                  <h2 className="text-3xl font-black mb-8 flex items-center gap-4 uppercase tracking-tighter">
                    <div className="bg-indigo-600 p-2.5 rounded-2xl shadow-lg">
                      <span className="text-white text-xl font-serif">Σ</span>
                    </div>
                    Tier 1: Basic Experiments
                  </h2>
                  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {pre100Exps.map(id => {
                      const summary = results?.experiments[id];
                      return (
                        <ExperimentCard 
                          key={id}
                          id={id}
                          name={id.replace(/_/g, ' ')}
                          level={plan.experimentLevels[id] || 0}
                          maxLevels={experimentMaxLevels[id]}
                          description={experimentDescriptions[id]}
                          currentValue={summary?.currentLevelValue || '0%'}
                          nextCost={summary?.nextLevelCost || '0'}
                          isMaxed={summary?.maxLevel || false}
                          onLevelChange={(lvl) => updateExpLevel(id, lvl)}
                        />
                      );
                    })}
                  </div>
                </section>

                <section>
                  <h2 className="text-3xl font-black mb-8 flex items-center gap-4 uppercase tracking-tighter">
                    <div className="bg-emerald-600 p-2.5 rounded-2xl shadow-lg">
                      <span className="text-white text-xl font-serif">Ω</span>
                    </div>
                    Tier 2: Advanced Multipliers
                  </h2>
                  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {post100Exps.map(id => {
                      const summary = results?.experiments[id];
                      return (
                        <ExperimentCard 
                          key={id}
                          id={id}
                          name={id.replace(/_/g, ' ')}
                          level={plan.experimentLevels[id] || 0}
                          maxLevels={experimentMaxLevels[id]}
                          description={experimentDescriptions[id]}
                          currentValue={summary?.currentLevelValue || 'x1.0'}
                          nextCost={summary?.nextLevelCost || '0'}
                          isMaxed={summary?.maxLevel || false}
                          onLevelChange={(lvl) => updateExpLevel(id, lvl)}
                        />
                      );
                    })}
                  </div>
                </section>
              </div>
            )}

            {/* Legendaries Tab */}
            {activeTab === 'legendaries' && (
              <div className="animate-in fade-in slide-in-from-bottom-4 duration-500">
                <h2 className="text-3xl font-black mb-8 flex items-center gap-4 uppercase tracking-tighter">
                  <div className="bg-amber-600 p-2.5 rounded-2xl shadow-lg">
                    <span className="text-white text-xl font-serif">★</span>
                  </div>
                  Planned Legendaries
                </h2>
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                  {legendaries.map(leg => (
                    <LegendaryCard 
                      key={leg.id}
                      id={leg.id}
                      name={leg.name}
                      count={plan.legendaryCounts[leg.id] || 0}
                      bonus={leg.bonus}
                      subsequent={leg.subsequent}
                      max={leg.max}
                      onCountChange={(cnt) => updateLegendaryCount(leg.id, cnt)}
                    />
                  ))}
                </div>
              </div>
            )}

            {/* Resources Tab */}
            {activeTab === 'resources' && (
              <div className="space-y-12 animate-in fade-in slide-in-from-bottom-4 duration-500">
                <section>
                  <h2 className="text-3xl font-black mb-8 flex items-center gap-4 uppercase tracking-tighter">
                    <div className="bg-sky-600 p-2.5 rounded-2xl shadow-lg">
                      <span className="text-white text-xl font-serif">◈</span>
                    </div>
                    Possessed Runes
                  </h2>
                  <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-6 gap-4">
                    {runeTypes.map(rune => (
                      <div key={rune} className="bg-slate-800/40 border border-slate-700/50 rounded-2xl p-4 flex flex-col items-center gap-3">
                        <img src={`/assets/images/${rune}_rune.png`} className="w-8 h-8" alt={rune}/>
                        <input 
                          type="number" 
                          value={plan.possessedRunes[rune] || 0}
                          onChange={(e) => updatePossessedRune(rune, parseInt(e.target.value) || 0)}
                          className="w-full bg-slate-950 border border-slate-800 rounded-xl text-center text-sm font-bold text-sky-400 py-2 focus:ring-1 focus:ring-sky-500 outline-none"
                        />
                      </div>
                    ))}
                  </div>
                </section>

                <section>
                  <h2 className="text-3xl font-black mb-8 flex items-center gap-4 uppercase tracking-tighter">
                    <div className="bg-rose-600 p-2.5 rounded-2xl shadow-lg">
                      <span className="text-white text-xl font-serif">♥</span>
                    </div>
                    Legendaries On-Board
                  </h2>
                  <div className="grid grid-cols-2 sm:grid-cols-4 md:grid-cols-6 gap-4">
                    {legendaries.map(leg => (
                      <div key={leg.id} className="bg-slate-800/40 border border-slate-700/50 rounded-2xl p-4 flex flex-col items-center gap-3">
                        <img src={`/assets/images/${leg.id}.webp`} className="w-8 h-8 object-contain opacity-60" alt={leg.name}/>
                        <input 
                          type="number" 
                          value={plan.possessedLegendaries[leg.id] || 0}
                          onChange={(e) => updatePossessedLegendary(leg.id, parseInt(e.target.value) || 0)}
                          className="w-full bg-slate-950 border border-slate-800 rounded-xl text-center text-sm font-bold text-rose-400 py-2 focus:ring-1 focus:ring-rose-500 outline-none"
                        />
                      </div>
                    ))}
                  </div>
                </section>
              </div>
            )}

            {/* Notes Section */}
            <section className="bg-slate-800/10 border border-slate-800 rounded-3xl p-8">
               <h3 className="text-[10px] font-black text-slate-500 uppercase tracking-widest mb-4">Planner Notes</h3>
               <textarea 
                value={plan.notes}
                onChange={(e) => updatePlan({ notes: e.target.value })}
                placeholder="Write your prestige strategy here..."
                className="w-full h-32 bg-slate-950/50 border border-slate-800 rounded-2xl p-4 text-sm text-slate-300 focus:ring-1 focus:ring-indigo-500 outline-none resize-none"
               />
            </section>
            
            <div className="flex justify-end gap-3 pb-20">
              <button onClick={handleReset} className="px-8 py-3 bg-slate-950 border border-slate-800 rounded-2xl text-[10px] font-black uppercase tracking-widest text-slate-500 hover:text-red-400 transition-all">Reset All</button>
              <button onClick={handleSave} className="px-12 py-3 bg-indigo-600 hover:bg-indigo-500 rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-indigo-600/30 transition-all">Save Simulation</button>
            </div>
          </div>
        </main>

        <aside className="w-[400px] border-l border-slate-800 bg-slate-950 p-8 flex flex-col gap-8 overflow-y-auto custom-scrollbar">
          <div className="space-y-8">
            <header>
               <h3 className="text-[10px] font-black text-slate-500 uppercase tracking-[0.2em] mb-4">Simulation Result</h3>
               <div className="flex items-center gap-5">
                  <div className="bg-indigo-600 p-4 rounded-3xl shadow-xl shadow-indigo-600/20">
                    <img src="/assets/images/time_shard.png" className="w-10 h-10" alt="Shards"/>
                  </div>
                  <div>
                    <div className="text-4xl font-black text-white leading-none mb-1">
                      {results?.totalShards.toLocaleString() || 0}
                    </div>
                    <div className="text-[10px] font-black text-indigo-400 uppercase tracking-widest">Calculated Shards</div>
                  </div>
               </div>
            </header>

            <div className="space-y-4 bg-slate-900/40 p-6 rounded-3xl border border-slate-800/50">
              <div className="flex justify-between items-center text-xs">
                <span className="text-slate-500 font-black uppercase tracking-tighter">Base Shards</span>
                <span className="font-mono text-slate-300">{(results?.baseShards || 0).toLocaleString()}</span>
              </div>
              <div className="flex justify-between items-center text-xs">
                <span className="text-slate-500 font-black uppercase tracking-tighter">Feat Multiplier</span>
                <span className="font-mono text-emerald-400">x{(results?.featMultiplier || 1).toFixed(2)}</span>
              </div>
              <div className="flex justify-between items-center text-xs">
                <span className="text-slate-500 font-black uppercase tracking-tighter">Legend Multiplier</span>
                <span className="font-mono text-amber-400">x{(results?.legendMultiplier || 1).toFixed(2)}</span>
              </div>
              <div className="flex justify-between items-center text-xs">
                <span className="text-slate-500 font-black uppercase tracking-tighter">Other Multiplier</span>
                <span className="font-mono text-indigo-400">x{(results?.otherMultiplier || 1).toFixed(2)}</span>
              </div>
              
              <div className="pt-4 border-t border-slate-800/50 space-y-3">
                <div className="flex justify-between items-center text-sm">
                  <span className="text-slate-500 font-bold">Planned Spend:</span>
                  <span className="font-mono text-red-400">-{results?.experimentCost.toLocaleString() || 0}</span>
                </div>
                <div className="flex justify-between items-end">
                  <span className="text-[10px] font-black text-white uppercase tracking-widest mb-1">Net Surplus</span>
                  <span className={`text-3xl font-black tabular-nums ${(results?.remaining || 0) >= 0 ? 'text-emerald-400' : 'text-red-500'}`}>
                    {results?.remaining.toLocaleString() || 0}
                  </span>
                </div>
              </div>
            </div>

            <div className="space-y-4">
              <h4 className="text-[10px] font-black text-slate-500 uppercase tracking-[0.2em]">Rune Deficiencies</h4>
              <div className="grid grid-cols-1 gap-2">
                {results && runeTypes.map(rune => {
                  const count = results.runeNeeded[rune];
                  return count > 0 ? (
                    <div key={rune} className="bg-slate-900/60 border border-slate-800 p-4 rounded-2xl flex justify-between items-center group hover:border-slate-700 transition-all">
                      <div className="flex items-center gap-3">
                        <div className="p-2 bg-slate-950 rounded-lg border border-slate-800">
                          <img src={`/assets/images/${rune}_rune.png`} className="w-5 h-5" alt={rune}/>
                        </div>
                        <span className="text-[10px] font-black text-slate-400 uppercase tracking-widest">{rune}</span>
                      </div>
                      <span className="text-lg font-black font-mono text-red-500">{count.toLocaleString()}</span>
                    </div>
                  ) : null;
                })}
                {results && Object.values(results.runeNeeded).every(c => c <= 0) && (
                  <div className="p-8 text-center border-2 border-dashed border-slate-800 rounded-3xl">
                    <div className="text-[10px] font-black text-slate-600 uppercase tracking-widest mb-2">Requirement Met</div>
                    <div className="text-emerald-500 text-sm font-bold uppercase">All Runes Available</div>
                  </div>
                )}
              </div>
            </div>
          </div>
        </aside>
      </div>
    </div>
  );
};

export default App;
