import React from 'react';
import { Activity, Zap, AlertTriangle, CheckCircle } from 'lucide-react';

const SignalsPanel: React.FC = () => {
    return (
        <div className="bg-surface p-6 rounded-lg shadow-lg mt-6">
            <h3 className="text-xl font-bold mb-4 flex items-center gap-2">
                <Zap className="text-accent" />
                Live Signals
            </h3>

            <div className="space-y-3">
                <div className="flex items-center justify-between p-3 bg-background rounded border-l-4 border-emerald-500">
                    <div className="flex items-center gap-2">
                        <CheckCircle size={18} className="text-emerald-500" />
                        <span className="font-bold">BTC/USDT Buy Signal</span>
                    </div>
                    <span className="text-xs text-secondary">2 min ago</span>
                </div>
                <div className="flex items-center justify-between p-3 bg-background rounded border-l-4 border-red-500">
                    <div className="flex items-center gap-2">
                        <AlertTriangle size={18} className="text-red-500" />
                        <span className="font-bold">ETH/USDT Resistance Hit</span>
                    </div>
                    <span className="text-xs text-secondary">15 min ago</span>
                </div>
                <div className="flex items-center justify-between p-3 bg-background rounded border-l-4 border-yellow-500">
                    <div className="flex items-center gap-2">
                        <Activity size={18} className="text-yellow-500" />
                        <span className="font-bold">SOL/USDT Volatility Alert</span>
                    </div>
                    <span className="text-xs text-secondary">1h ago</span>
                </div>
            </div>
        </div>
    );
};

export default SignalsPanel;
