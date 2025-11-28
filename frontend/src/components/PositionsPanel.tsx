import React, { useState } from 'react';
import { ArrowUpCircle, ArrowDownCircle, XCircle, Clock, CheckCircle2 } from 'lucide-react';

type Tab = 'positions' | 'open-orders' | 'history';

const PositionsPanel: React.FC = () => {
    const [activeTab, setActiveTab] = useState<Tab>('positions');

    // Mock Data
    const positions = [
        { id: 1, symbol: 'BTCUSDT', side: 'LONG', size: 0.5, entryPrice: 49500.00, markPrice: 50120.50, pnl: 310.25, pnlPercent: 1.25 },
        { id: 2, symbol: 'ETHUSDT', side: 'SHORT', size: 5.0, entryPrice: 3200.00, markPrice: 3150.00, pnl: 250.00, pnlPercent: 1.56 },
        { id: 3, symbol: 'SOLUSDT', side: 'LONG', size: 100, entryPrice: 145.00, markPrice: 142.50, pnl: -250.00, pnlPercent: -1.72 },
    ];

    const openOrders = [
        { id: 101, symbol: 'XRPUSDT', side: 'BUY', type: 'LIMIT', size: 1000, price: 0.5500, time: '10:30:00' },
        { id: 102, symbol: 'BTCUSDT', side: 'SELL', type: 'STOP-LOSS', size: 0.5, price: 49000.00, time: '10:35:00' },
    ];

    const history = [
        { id: 201, symbol: 'ADAUSDT', side: 'SELL', size: 5000, price: 0.4500, realizedPnl: 120.50, time: '09:15:00' },
        { id: 202, symbol: 'DOTUSDT', side: 'BUY', size: 200, price: 7.50, realizedPnl: -15.00, time: '08:45:00' },
    ];

    return (
        <div className="bg-surface p-6 rounded-lg shadow-lg mt-6">
            <div className="flex items-center gap-6 border-b border-secondary/20 mb-6">
                <button
                    onClick={() => setActiveTab('positions')}
                    className={`pb-3 font-bold text-sm transition-colors ${activeTab === 'positions' ? 'text-accent border-b-2 border-accent' : 'text-secondary hover:text-primary'}`}
                >
                    Positions ({positions.length})
                </button>
                <button
                    onClick={() => setActiveTab('open-orders')}
                    className={`pb-3 font-bold text-sm transition-colors ${activeTab === 'open-orders' ? 'text-accent border-b-2 border-accent' : 'text-secondary hover:text-primary'}`}
                >
                    Open Orders ({openOrders.length})
                </button>
                <button
                    onClick={() => setActiveTab('history')}
                    className={`pb-3 font-bold text-sm transition-colors ${activeTab === 'history' ? 'text-accent border-b-2 border-accent' : 'text-secondary hover:text-primary'}`}
                >
                    Order History
                </button>
            </div>

            <div className="overflow-x-auto">
                {activeTab === 'positions' && (
                    <table className="w-full text-left text-sm">
                        <thead>
                            <tr className="text-secondary border-b border-secondary/20">
                                <th className="pb-3 pl-2">Symbol</th>
                                <th className="pb-3">Side</th>
                                <th className="pb-3">Size</th>
                                <th className="pb-3">Entry Price</th>
                                <th className="pb-3">Mark Price</th>
                                <th className="pb-3">PnL (ROE%)</th>
                                <th className="pb-3 text-right pr-2">Action</th>
                            </tr>
                        </thead>
                        <tbody>
                            {positions.map((pos) => (
                                <tr key={pos.id} className="border-b border-secondary/10 hover:bg-background/50 transition-colors">
                                    <td className="py-4 pl-2 font-bold">{pos.symbol}</td>
                                    <td className={`py-4 font-bold ${pos.side === 'LONG' ? 'text-emerald-500' : 'text-red-500'}`}>
                                        {pos.side}
                                    </td>
                                    <td className="py-4">{pos.size}</td>
                                    <td className="py-4">${pos.entryPrice.toFixed(2)}</td>
                                    <td className="py-4">${pos.markPrice.toFixed(2)}</td>
                                    <td className={`py-4 font-bold ${pos.pnl >= 0 ? 'text-emerald-500' : 'text-red-500'}`}>
                                        {pos.pnl >= 0 ? '+' : ''}{pos.pnl.toFixed(2)} ({pos.pnlPercent}%)
                                    </td>
                                    <td className="py-4 text-right pr-2">
                                        <button className="text-secondary hover:text-red-500 transition-colors">
                                            <XCircle size={18} />
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                )}

                {activeTab === 'open-orders' && (
                    <table className="w-full text-left text-sm">
                        <thead>
                            <tr className="text-secondary border-b border-secondary/20">
                                <th className="pb-3 pl-2">Time</th>
                                <th className="pb-3">Symbol</th>
                                <th className="pb-3">Type</th>
                                <th className="pb-3">Side</th>
                                <th className="pb-3">Price</th>
                                <th className="pb-3">Size</th>
                                <th className="pb-3 text-right pr-2">Action</th>
                            </tr>
                        </thead>
                        <tbody>
                            {openOrders.map((order) => (
                                <tr key={order.id} className="border-b border-secondary/10 hover:bg-background/50 transition-colors">
                                    <td className="py-4 pl-2 text-secondary">{order.time}</td>
                                    <td className="py-4 font-bold">{order.symbol}</td>
                                    <td className="py-4">{order.type}</td>
                                    <td className={`py-4 font-bold ${order.side === 'BUY' ? 'text-emerald-500' : 'text-red-500'}`}>
                                        {order.side}
                                    </td>
                                    <td className="py-4">${order.price.toFixed(4)}</td>
                                    <td className="py-4">{order.size}</td>
                                    <td className="py-4 text-right pr-2">
                                        <button className="text-secondary hover:text-red-500 transition-colors">
                                            Cancel
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                )}

                {activeTab === 'history' && (
                    <table className="w-full text-left text-sm">
                        <thead>
                            <tr className="text-secondary border-b border-secondary/20">
                                <th className="pb-3 pl-2">Time</th>
                                <th className="pb-3">Symbol</th>
                                <th className="pb-3">Side</th>
                                <th className="pb-3">Price</th>
                                <th className="pb-3">Size</th>
                                <th className="pb-3">Realized PnL</th>
                            </tr>
                        </thead>
                        <tbody>
                            {history.map((trade) => (
                                <tr key={trade.id} className="border-b border-secondary/10 hover:bg-background/50 transition-colors">
                                    <td className="py-4 pl-2 text-secondary">{trade.time}</td>
                                    <td className="py-4 font-bold">{trade.symbol}</td>
                                    <td className={`py-4 font-bold ${trade.side === 'BUY' ? 'text-emerald-500' : 'text-red-500'}`}>
                                        {trade.side}
                                    </td>
                                    <td className="py-4">${trade.price.toFixed(4)}</td>
                                    <td className="py-4">{trade.size}</td>
                                    <td className={`py-4 font-bold ${trade.realizedPnl >= 0 ? 'text-emerald-500' : 'text-red-500'}`}>
                                        {trade.realizedPnl >= 0 ? '+' : ''}{trade.realizedPnl.toFixed(2)}
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                )}
            </div>
        </div>
    );
};

export default PositionsPanel;
