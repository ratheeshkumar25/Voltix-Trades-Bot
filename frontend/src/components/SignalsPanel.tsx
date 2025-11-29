import React, { useState } from 'react';
import { Zap, AlertTriangle, CheckCircle, TrendingUp, Clock, Target, DollarSign, ChevronDown, ChevronUp } from 'lucide-react';

interface Signal {
    id: number;
    symbol: string;
    type: 'BUY' | 'SELL';
    entryPrice: number;
    targetPrice: number;
    stopLoss: number;
    expectedProfit: number;
    timeFrame: string;
    signalStrength: 'HIGH' | 'MEDIUM' | 'LOW';
    timestamp: string;
    validUntil: string;
}

const SignalsPanel: React.FC = () => {
    const [expandedSignal, setExpandedSignal] = useState<number | null>(null);

    // Mock signal data with comprehensive trading information
    const signals: Signal[] = [
        {
            id: 1,
            symbol: 'BTC/USDT',
            type: 'BUY',
            entryPrice: 50100,
            targetPrice: 52000,
            stopLoss: 49200,
            expectedProfit: 3.8,
            timeFrame: '2-4 hours',
            signalStrength: 'HIGH',
            timestamp: '2 min ago',
            validUntil: '30 min'
        },
        {
            id: 2,
            symbol: 'ETH/USDT',
            type: 'SELL',
            entryPrice: 3150,
            targetPrice: 3050,
            stopLoss: 3200,
            expectedProfit: 3.2,
            timeFrame: '1-2 hours',
            signalStrength: 'MEDIUM',
            timestamp: '15 min ago',
            validUntil: '15 min'
        },
        {
            id: 3,
            symbol: 'SOL/USDT',
            type: 'BUY',
            entryPrice: 142.50,
            targetPrice: 148.00,
            stopLoss: 139.00,
            expectedProfit: 3.9,
            timeFrame: '3-6 hours',
            signalStrength: 'HIGH',
            timestamp: '1h ago',
            validUntil: '2h 30min'
        }
    ];

    const getSignalColor = (type: 'BUY' | 'SELL') => {
        return type === 'BUY' ? 'border-emerald-500' : 'border-red-500';
    };

    const getSignalIcon = (type: 'BUY' | 'SELL') => {
        return type === 'BUY' ? (
            <CheckCircle size={18} className="text-emerald-500" />
        ) : (
            <AlertTriangle size={18} className="text-red-500" />
        );
    };

    const getStrengthBadge = (strength: 'HIGH' | 'MEDIUM' | 'LOW') => {
        const colors = {
            HIGH: 'bg-emerald-500/20 text-emerald-500',
            MEDIUM: 'bg-yellow-500/20 text-yellow-500',
            LOW: 'bg-gray-500/20 text-gray-500'
        };
        return (
            <span className={`${colors[strength]} text-xs px-2 py-1 rounded font-bold`}>
                {strength}
            </span>
        );
    };

    const toggleSignal = (id: number) => {
        setExpandedSignal(expandedSignal === id ? null : id);
    };

    return (
        <div className="bg-surface p-6 rounded-lg shadow-lg mt-6">
            <h3 className="text-xl font-bold mb-4 flex items-center gap-2">
                <Zap className="text-accent" />
                Live Signals
            </h3>

            <div className="space-y-3">
                {signals.map((signal) => (
                    <div
                        key={signal.id}
                        className={`bg-background rounded border-l-4 ${getSignalColor(signal.type)} overflow-hidden transition-all`}
                    >
                        {/* Signal Header */}
                        <div
                            className="flex items-center justify-between p-3 cursor-pointer hover:bg-surface/30 transition-colors"
                            onClick={() => toggleSignal(signal.id)}
                        >
                            <div className="flex items-center gap-2 flex-1">
                                {getSignalIcon(signal.type)}
                                <div>
                                    <span className="font-bold">{signal.symbol} {signal.type} Signal</span>
                                    <div className="flex items-center gap-2 mt-1">
                                        {getStrengthBadge(signal.signalStrength)}
                                        <span className="text-xs text-secondary flex items-center gap-1">
                                            <Clock size={12} />
                                            {signal.timestamp}
                                        </span>
                                    </div>
                                </div>
                            </div>
                            <div className="flex items-center gap-2">
                                <div className="text-right mr-2">
                                    <div className={`text-sm font-bold ${signal.type === 'BUY' ? 'text-emerald-500' : 'text-red-500'}`}>
                                        +{signal.expectedProfit}%
                                    </div>
                                    <div className="text-xs text-secondary">{signal.timeFrame}</div>
                                </div>
                                {expandedSignal === signal.id ? (
                                    <ChevronUp size={20} className="text-secondary" />
                                ) : (
                                    <ChevronDown size={20} className="text-secondary" />
                                )}
                            </div>
                        </div>

                        {/* Expanded Details */}
                        {expandedSignal === signal.id && (
                            <div className="px-3 pb-3 border-t border-secondary/20">
                                <div className="grid grid-cols-2 gap-3 mt-3">
                                    {/* Entry Price */}
                                    <div className="p-2 bg-surface/50 rounded">
                                        <div className="text-xs text-secondary mb-1 flex items-center gap-1">
                                            <DollarSign size={12} />
                                            Entry Price
                                        </div>
                                        <div className="font-bold text-primary">${signal.entryPrice.toLocaleString()}</div>
                                    </div>

                                    {/* Target Price */}
                                    <div className="p-2 bg-surface/50 rounded">
                                        <div className="text-xs text-secondary mb-1 flex items-center gap-1">
                                            <Target size={12} />
                                            Target Price
                                        </div>
                                        <div className={`font-bold ${signal.type === 'BUY' ? 'text-emerald-500' : 'text-red-500'}`}>
                                            ${signal.targetPrice.toLocaleString()}
                                        </div>
                                    </div>

                                    {/* Stop Loss */}
                                    <div className="p-2 bg-surface/50 rounded">
                                        <div className="text-xs text-secondary mb-1">Stop Loss</div>
                                        <div className="font-bold text-red-500">${signal.stopLoss.toLocaleString()}</div>
                                    </div>

                                    {/* Expected Profit */}
                                    <div className="p-2 bg-surface/50 rounded">
                                        <div className="text-xs text-secondary mb-1 flex items-center gap-1">
                                            <TrendingUp size={12} />
                                            Expected Profit
                                        </div>
                                        <div className="font-bold text-emerald-500">+{signal.expectedProfit}%</div>
                                    </div>
                                </div>

                                {/* Trade Timeline Section */}
                                <div className="mt-3 p-3 bg-accent/10 rounded border border-accent/30">
                                    <div className="text-xs font-bold text-accent mb-2">📊 Trade Timeline & Profit Strategy</div>
                                    <div className="text-xs text-gray-300 space-y-1">
                                        <div className="flex items-start gap-2">
                                            <span className="text-accent">•</span>
                                            <span>
                                                <strong>Entry:</strong> Execute {signal.type} at ${signal.entryPrice.toLocaleString()}
                                                {signal.validUntil && <span className="text-yellow-500"> (Valid for {signal.validUntil})</span>}
                                            </span>
                                        </div>
                                        <div className="flex items-start gap-2">
                                            <span className="text-accent">•</span>
                                            <span>
                                                <strong>Hold Period:</strong> {signal.timeFrame} to reach target of ${signal.targetPrice.toLocaleString()}
                                            </span>
                                        </div>
                                        <div className="flex items-start gap-2">
                                            <span className="text-accent">•</span>
                                            <span>
                                                <strong>Exit Strategy:</strong> Take profit at ${signal.targetPrice.toLocaleString()} (+{signal.expectedProfit}%) or cut loss at ${signal.stopLoss.toLocaleString()}
                                            </span>
                                        </div>
                                        <div className="flex items-start gap-2">
                                            <span className="text-accent">•</span>
                                            <span>
                                                <strong>Risk/Reward:</strong> Potential gain of {signal.expectedProfit}% over {signal.timeFrame}
                                            </span>
                                        </div>
                                    </div>
                                </div>

                                {/* Action Button */}
                                <button className={`w-full mt-3 py-2 rounded font-bold transition-colors ${signal.type === 'BUY'
                                    ? 'bg-emerald-500 hover:bg-emerald-600 text-white'
                                    : 'bg-red-500 hover:bg-red-600 text-white'
                                    }`}>
                                    Execute {signal.type} Trade
                                </button>
                            </div>
                        )}
                    </div>
                ))}
            </div>
        </div>
    );
};

export default SignalsPanel;
