import React from 'react';
import type { ExperimentID } from '../types/game';

interface Props {
  id: ExperimentID;
  name: string;
  level: number;
  maxLevels: number;
  description: string;
  currentValue: string;
  nextCost: string;
  isMaxed: boolean;
  onLevelChange: (level: number) => void;
}

export const ExperimentCard: React.FC<Props> = ({ 
  id, name, level, maxLevels, description, currentValue, nextCost, isMaxed, onLevelChange 
}) => {
  return (
    <div className="bg-slate-800/40 border border-slate-700/50 rounded-3xl p-6 hover:border-slate-600 transition-all group backdrop-blur-sm">
      <div className="flex items-center gap-4 mb-4">
        <div className="w-16 h-16 rounded-2xl bg-slate-900 border border-slate-700/50 p-2 flex items-center justify-center shadow-inner group-hover:scale-105 transition-transform shrink-0">
          <img src={`/assets/images/${id}.png`} alt={name} className="w-12 h-12 object-contain" onError={(e) => { (e.target as HTMLImageElement).src = '/assets/images/time_shard.png'; (e.target as HTMLImageElement).style.opacity = '0.2'; }} />
        </div>
        <div className="flex-1 min-w-0">
          <div className="flex justify-between items-center mb-1">
            <h4 className="text-lg font-black text-white leading-tight uppercase truncate pr-4">{name}</h4>
            <div className="bg-slate-900 px-3 py-1.5 rounded-xl border border-slate-800 shrink-0">
              <span className="text-[10px] font-black text-slate-500 uppercase tracking-tighter">
                LVL {level}
              </span>
            </div>
          </div>
        </div>
      </div>

      <p className="text-xs text-slate-400 mb-4 h-10 overflow-hidden line-clamp-2">{description}</p>

      <div className="relative mb-6">
        <input 
          type="range"
          min="0"
          max={maxLevels}
          value={level}
          onChange={(e) => onLevelChange(parseInt(e.target.value))}
          className="w-full h-2 bg-slate-900 rounded-full appearance-none cursor-pointer accent-indigo-500 focus:ring-2 focus:ring-indigo-500/20"
        />
        <div className="flex justify-between mt-3 px-1">
          {Array.from({ length: maxLevels + 1 }).map((_, i) => (
            <div 
              key={i}
              className={`w-1 h-1 rounded-full transition-colors duration-300 ${level >= i ? 'bg-indigo-500 shadow-[0_0_8px_rgba(99,102,241,0.6)]' : 'bg-slate-700'}`}
            />
          ))}
        </div>
      </div>

      <div className="flex justify-between items-center bg-slate-900/50 p-4 rounded-2xl border border-slate-800/50">
        <div>
          <p className="text-[9px] font-black text-slate-500 uppercase tracking-widest mb-1">Current Effect</p>
          <span className="text-sm font-bold text-indigo-400 font-mono">{currentValue}</span>
        </div>
        <div className="text-right">
          <p className="text-[9px] font-black text-slate-500 uppercase tracking-widest mb-1">Next Tier</p>
          <div className="flex items-center gap-2 justify-end">
             <span className={`text-sm font-bold font-mono ${isMaxed ? 'text-emerald-500' : 'text-white'}`}>
                {isMaxed ? 'MAXED' : nextCost}
              </span>
              {!isMaxed && <img src="/assets/images/time_shard.png" className="w-4 h-4 opacity-80" alt="shards" />}
          </div>
        </div>
      </div>
    </div>
  );
};
