import React from 'react';
import type { LegendaryID } from '../types/game';

interface Props {
  id: LegendaryID;
  name: string;
  count: number;
  bonus: string;
  subsequent: string;
  max?: number;
  onCountChange: (count: number) => void;
}

export const LegendaryCard: React.FC<Props> = ({ 
  id, name, count, bonus, subsequent, max, onCountChange 
}) => {
  return (
    <div className="bg-slate-800/40 border border-slate-700/50 rounded-2xl p-4 hover:border-slate-600 transition-all group flex items-center gap-4">
      <div className="w-12 h-12 rounded-xl bg-slate-900 border border-slate-700/50 p-1 flex items-center justify-center shrink-0">
        <img src={`/assets/images/${id}.webp`} alt={name} className="w-10 h-10 object-contain" onError={(e) => { (e.target as HTMLImageElement).src = '/assets/images/time_shard.png'; (e.target as HTMLImageElement).style.opacity = '0.2'; }} />
      </div>
      
      <div className="flex-1 min-w-0">
        <h4 className="text-sm font-black text-white uppercase truncate">{name}</h4>
        <div className="text-[10px] text-slate-500 font-mono">
          {bonus} / {subsequent}
        </div>
      </div>

      <div className="flex items-center bg-slate-900 rounded-xl border border-slate-800 p-1">
        <button 
          onClick={() => onCountChange(Math.max(0, count - 1))}
          className="w-8 h-8 flex items-center justify-center text-slate-400 hover:text-white hover:bg-slate-800 rounded-lg transition-colors"
        >
          -
        </button>
        <input 
          type="number" 
          value={count} 
          onChange={(e) => onCountChange(Math.max(0, parseInt(e.target.value) || 0))}
          className="w-8 text-center bg-transparent border-none text-sm font-bold text-indigo-400 focus:ring-0"
        />
        <button 
          onClick={() => onCountChange(max ? Math.min(max, count + 1) : count + 1)}
          className="w-8 h-8 flex items-center justify-center text-slate-400 hover:text-white hover:bg-slate-800 rounded-lg transition-colors"
        >
          +
        </button>
      </div>
    </div>
  );
};
