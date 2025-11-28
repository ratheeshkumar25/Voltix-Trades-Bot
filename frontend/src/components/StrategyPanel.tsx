import React from 'react';
import { Activity, Zap, AlertTriangle, CheckCircle } from 'lucide-react';

const StrategyPanel: React.FC = () => {
    return (
        <div className="bg-surface p-6 rounded-lg shadow-lg mt-6">
            <h3 className="text-xl font-bold mb-4 flex items-center gap-2">
                <Activity className="text-accent" />
                Active Trading Strategy
            </h3>

            <div>
                <div className="flex items-center justify-between mb-2">
                    <span className="text-secondary">Strategy Name</span>
                    <span className="font-bold text-primary">Hybrid Momentum & Reversal</span>
                </div>
                <div className="flex items-center justify-between mb-2">
                    <span className="text-secondary">Risk Level</span>
                    <span className="font-bold text-yellow-500">Moderate</span>
                </div>
                <div className="flex items-center justify-between mb-2">
                    <span className="text-secondary">Timeframe</span>
                    <span className="font-bold text-primary">15m / 1h / 4h</span>
                </div>

                <div className="mt-4 p-3 bg-background rounded border border-secondary/30">
                    <h4 className="font-bold mb-2 text-sm text-secondary">Strategy Description</h4>
                    <p className="text-sm text-gray-400">
                        Combines RSI divergence with MACD crossovers to identify trend reversals.
                        Uses Pivot Points for dynamic support and resistance levels to set take-profit and stop-loss targets.
                    </p>
                </div>
            </div>
        </div>
    );
};

export default StrategyPanel;
