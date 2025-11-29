import React from 'react';
import { Activity, Zap, TrendingUp, TrendingDown } from 'lucide-react';

const StrategyPanel: React.FC = () => {
    // Simulated indicator data - in production, this would come from real-time data
    const indicators = {
        pivotPoints: {
            status: 'active',
            resistance: 50500,
            support: 48900,
            current: 50120,
        },
        rsi: {
            status: 'active',
            value: 58.3,
            signal: 'neutral', // 'overbought', 'oversold', 'neutral'
        },
        macd: {
            status: 'active',
            signal: 'bullish', // 'bullish', 'bearish'
            histogram: 45.6,
        }
    };

    const getStatusBadge = (status: string) => {
        const colors = {
            active: 'bg-emerald-500',
            monitoring: 'bg-yellow-500',
            inactive: 'bg-gray-500'
        };
        return (
            <span className={`${colors[status as keyof typeof colors]} text-white text-xs px-2 py-1 rounded-full font-bold`}>
                {status.toUpperCase()}
            </span>
        );
    };

    const getRSIColor = (value: number) => {
        if (value >= 70) return 'text-red-500';
        if (value <= 30) return 'text-emerald-500';
        return 'text-yellow-500';
    };

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

                {/* Active Indicators Section */}
                <div className="mt-6 border-t border-secondary/30 pt-4">
                    <h4 className="font-bold mb-4 text-lg flex items-center gap-2">
                        <Zap className="text-accent" size={18} />
                        Active Indicators
                    </h4>

                    {/* Pivot Points */}
                    <div className="mb-4 p-3 bg-background rounded border border-secondary/20">
                        <div className="flex items-center justify-between mb-2">
                            <span className="font-bold text-sm">Pivot Points</span>
                            {getStatusBadge(indicators.pivotPoints.status)}
                        </div>
                        <div className="grid grid-cols-3 gap-2 text-xs">
                            <div>
                                <div className="text-secondary">Support</div>
                                <div className="font-bold text-emerald-500">${indicators.pivotPoints.support.toLocaleString()}</div>
                            </div>
                            <div>
                                <div className="text-secondary">Current</div>
                                <div className="font-bold text-primary">${indicators.pivotPoints.current.toLocaleString()}</div>
                            </div>
                            <div>
                                <div className="text-secondary">Resistance</div>
                                <div className="font-bold text-red-500">${indicators.pivotPoints.resistance.toLocaleString()}</div>
                            </div>
                        </div>
                    </div>

                    {/* RSI */}
                    <div className="mb-4 p-3 bg-background rounded border border-secondary/20">
                        <div className="flex items-center justify-between mb-2">
                            <span className="font-bold text-sm">RSI (Relative Strength Index)</span>
                            {getStatusBadge(indicators.rsi.status)}
                        </div>
                        <div className="flex items-center justify-between">
                            <div>
                                <div className="text-xs text-secondary">Current Value</div>
                                <div className={`text-2xl font-bold ${getRSIColor(indicators.rsi.value)}`}>
                                    {indicators.rsi.value}
                                </div>
                            </div>
                            <div className="text-right">
                                <div className="text-xs text-secondary mb-1">Signal</div>
                                <span className="inline-flex items-center gap-1 px-2 py-1 rounded bg-yellow-500/20 text-yellow-500 text-xs font-bold">
                                    {indicators.rsi.signal === 'overbought' && <TrendingDown size={14} />}
                                    {indicators.rsi.signal === 'oversold' && <TrendingUp size={14} />}
                                    {indicators.rsi.signal.toUpperCase()}
                                </span>
                            </div>
                        </div>
                        <div className="mt-2 text-xs text-gray-400">
                            {indicators.rsi.value >= 70 && 'Overbought - Potential Sell Signal'}
                            {indicators.rsi.value <= 30 && 'Oversold - Potential Buy Signal'}
                            {indicators.rsi.value > 30 && indicators.rsi.value < 70 && 'Neutral Range - Monitor for breakout'}
                        </div>
                    </div>

                    {/* MACD */}
                    <div className="mb-2 p-3 bg-background rounded border border-secondary/20">
                        <div className="flex items-center justify-between mb-2">
                            <span className="font-bold text-sm">MACD</span>
                            {getStatusBadge(indicators.macd.status)}
                        </div>
                        <div className="flex items-center justify-between">
                            <div>
                                <div className="text-xs text-secondary">Histogram</div>
                                <div className={`text-xl font-bold ${indicators.macd.histogram >= 0 ? 'text-emerald-500' : 'text-red-500'}`}>
                                    {indicators.macd.histogram > 0 ? '+' : ''}{indicators.macd.histogram}
                                </div>
                            </div>
                            <div className="text-right">
                                <div className="text-xs text-secondary mb-1">Trend Signal</div>
                                <span className={`inline-flex items-center gap-1 px-2 py-1 rounded text-xs font-bold ${indicators.macd.signal === 'bullish'
                                    ? 'bg-emerald-500/20 text-emerald-500'
                                    : 'bg-red-500/20 text-red-500'
                                    }`}>
                                    {indicators.macd.signal === 'bullish' ? <TrendingUp size={14} /> : <TrendingDown size={14} />}
                                    {indicators.macd.signal.toUpperCase()}
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default StrategyPanel;
