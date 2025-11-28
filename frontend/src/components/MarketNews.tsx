import React from 'react';
import { TrendingUp, TrendingDown, AlertCircle, Newspaper } from 'lucide-react';

const MarketNews: React.FC = () => {
    const newsItems = [
        { id: 1, title: 'Bitcoin Breaks $50K Resistance', time: '2 min ago', type: 'bullish' },
        { id: 2, title: 'Federal Reserve Announces Rate Decision', time: '15 min ago', type: 'neutral' },
        { id: 3, title: 'Ethereum 2.0 Upgrade Completed', time: '1h ago', type: 'bullish' },
        { id: 4, title: 'Major Exchange Reports Volume Surge', time: '2h ago', type: 'bullish' }
    ];

    const alerts = [
        { id: 1, symbol: 'BTC/USDT', action: 'BUY', price: '50,245', reason: 'RSI oversold + MACD crossover', time: '5 min ago' },
        { id: 2, symbol: 'ETH/USDT', action: 'SELL', price: '3,120', reason: 'Resistance level reached', time: '12 min ago' },
        { id: 3, symbol: 'SOL/USDT', action: 'BUY', price: '105.50', reason: 'Support bounce detected', time: '25 min ago' }
    ];

    return (
        <div className="space-y-6">
            {/* Market News Section */}
            <div className="bg-surface p-6 rounded-lg shadow-lg">
                <h3 className="text-xl font-bold mb-4 flex items-center gap-2">
                    <Newspaper className="text-accent" />
                    Market News
                </h3>
                <div className="space-y-3">
                    {newsItems.map((news) => (
                        <div
                            key={news.id}
                            className="flex justify-between items-start p-3 bg-background rounded border-l-4 border-blue-500 hover:bg-secondary/10 transition-colors"
                        >
                            <div className="flex-1">
                                <p className="font-medium text-sm">{news.title}</p>
                                <p className="text-xs text-secondary mt-1">{news.time}</p>
                            </div>
                            {news.type === 'bullish' && <TrendingUp size={16} className="text-green-500 ml-2" />}
                            {news.type === 'bearish' && <TrendingDown size={16} className="text-red-500 ml-2" />}
                        </div>
                    ))}
                </div>
            </div>

            {/* Trading Alerts Section */}
            <div className="bg-surface p-6 rounded-lg shadow-lg">
                <h3 className="text-xl font-bold mb-4 flex items-center gap-2">
                    <AlertCircle className="text-accent" />
                    Trading Alerts
                </h3>
                <div className="space-y-3">
                    {alerts.map((alert) => (
                        <div
                            key={alert.id}
                            className={`p-3 bg-background rounded border-l-4 ${alert.action === 'BUY' ? 'border-green-500' : 'border-red-500'
                                }`}
                        >
                            <div className="flex justify-between items-center mb-2">
                                <div className="flex items-center gap-2">
                                    <span className="font-bold">{alert.symbol}</span>
                                    <span className={`px-2 py-0.5 rounded text-xs font-bold ${alert.action === 'BUY' ? 'bg-green-500/20 text-green-500' : 'bg-red-500/20 text-red-500'
                                        }`}>
                                        {alert.action}
                                    </span>
                                </div>
                                <span className="text-accent font-bold">${alert.price}</span>
                            </div>
                            <p className="text-xs text-gray-400 mb-1">{alert.reason}</p>
                            <p className="text-xs text-secondary">{alert.time}</p>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default MarketNews;
