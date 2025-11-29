import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { AdvancedRealTimeChart } from "react-ts-tradingview-widgets";
import { Wallet, ArrowRightLeft, Coins, Shield, Maximize2, Minimize2 } from 'lucide-react';
import TradePanel from './TradePanel';
import StrategyPanel from './StrategyPanel';
import PositionsPanel from './PositionsPanel';
import SignalsPanel from './SignalsPanel';
import MarketNews from './MarketNews';
import logo from '../assets/voltix-logo.jpg';

type AccountType = 'metatrader' | 'binance' | 'ctrader' | 'gmail';

interface DashboardProps {
    accountType: AccountType;
}

const Dashboard: React.FC<DashboardProps> = ({ accountType }) => {
    const [exchange, setExchange] = useState('binance');
    const [symbol, setSymbol] = useState('BTCUSDT');
    const [balance, setBalance] = useState(0);
    const [isChartFullscreen, setIsChartFullscreen] = useState(false);

    const isViewOnly = accountType === 'gmail';

    useEffect(() => {
        const fetchBalance = async () => {
            setBalance(10000); // Static balance
        };
        fetchBalance();
    }, [exchange]);

    // Handle ESC key to exit fullscreen
    useEffect(() => {
        const handleKeyDown = (e: KeyboardEvent) => {
            if (e.key === 'Escape' && isChartFullscreen) {
                setIsChartFullscreen(false);
            }
        };

        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
    }, [isChartFullscreen]);

    const toggleChartFullscreen = () => {
        setIsChartFullscreen(!isChartFullscreen);
    };

    const getAccountBadge = () => {
        const badges = {
            metatrader: { label: 'MetaTrader', color: 'bg-blue-500' },
            binance: { label: 'Binance', color: 'bg-yellow-500' },
            ctrader: { label: 'cTrader', color: 'bg-green-500' },
            gmail: { label: 'View Only', color: 'bg-gray-500' }
        };
        const badge = badges[accountType];
        return (
            <div className={`flex items-center gap-2 ${badge.color} px-4 py-2 rounded-lg text-white`}>
                <Shield size={18} />
                <span className="font-bold">{badge.label}</span>
            </div>
        );
    };

    return (
        <div className="min-h-screen bg-background p-6">
            <header className="flex justify-between items-center mb-8">
                <h1 className="text-3xl font-bold text-primary flex items-center gap-2">
                    <img src={logo} alt="Voltix Logo" className="h-10 w-auto rounded" />
                    Voltix Dashboard
                </h1>
                <div className="flex items-center gap-4">
                    {!isViewOnly && (
                        <div className="flex items-center gap-2 bg-surface px-4 py-2 rounded-lg">
                            <Wallet size={18} className="text-secondary" />
                            <span className="font-mono text-lg">${balance.toFixed(2)}</span>
                        </div>
                    )}
                    {getAccountBadge()}
                </div>
            </header>

            {!isViewOnly && (
                <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
                    <div className="flex items-center gap-2 bg-surface px-4 py-2 rounded-lg">
                        <ArrowRightLeft size={18} className="text-secondary" />
                        <label className="text-sm text-secondary">Exchange:</label>
                        <select
                            value={exchange}
                            onChange={(e) => setExchange(e.target.value)}
                            className="bg-background px-2 py-1 rounded border border-secondary focus:border-primary outline-none"
                        >
                            <option value="binance">Binance</option>
                            <option value="metatrader">MetaTrader</option>
                            <option value="ctrader">cTrader</option>
                        </select>
                    </div>
                    <div className="flex items-center gap-2 bg-surface px-4 py-2 rounded-lg">
                        <Coins size={18} className="text-secondary" />
                        <label className="text-sm text-secondary">Symbol:</label>
                        <select
                            value={symbol}
                            onChange={(e) => setSymbol(e.target.value)}
                            className="bg-background px-2 py-1 rounded border border-secondary focus:border-primary outline-none"
                        >
                            <option value="BTCUSDT">BTC/USDT</option>
                            <option value="ETHUSDT">ETH/USDT</option>
                            <option value="SOLUSDT">SOL/USDT</option>
                            <option value="XRPUSDT">XRP/USDT</option>
                        </select>
                    </div>
                </div>
            )}

            {/* Fullscreen Chart Overlay */}
            {isChartFullscreen && (
                <div className="fixed inset-0 z-50 bg-background p-6">
                    <div className="h-full flex flex-col">
                        <div className="flex justify-between items-center mb-4">
                            <h3 className="text-2xl font-bold">Market Overview - {symbol} ({exchange.toUpperCase()})</h3>
                            <button
                                onClick={toggleChartFullscreen}
                                className="flex items-center gap-2 bg-surface hover:bg-surface/80 px-4 py-2 rounded-lg transition-colors"
                            >
                                <Minimize2 size={20} />
                                <span>Exit Fullscreen</span>
                            </button>
                        </div>
                        <div className="flex-1">
                            <AdvancedRealTimeChart
                                theme="dark"
                                symbol={symbol}
                                autosize
                                hide_side_toolbar={false}
                                allow_symbol_change={true}
                                studies={[
                                    "PivotPointsStandard@tv-basicstudies",
                                    "RSI@tv-basicstudies",
                                    "MACD@tv-basicstudies"
                                ]}
                            />
                        </div>
                    </div>
                </div>
            )}

            <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                <div className="lg:col-span-2 bg-surface p-6 rounded-lg shadow-lg">
                    <div className="flex justify-between items-center mb-4">
                        <h3 className="text-xl font-bold">Market Overview ({exchange.toUpperCase()})</h3>
                        <button
                            onClick={toggleChartFullscreen}
                            className="flex items-center gap-2 bg-background hover:bg-background/80 px-3 py-2 rounded-lg transition-colors"
                            title="Toggle Fullscreen"
                        >
                            <Maximize2 size={18} />
                            <span className="text-sm">Fullscreen</span>
                        </button>
                    </div>
                    <div className="h-[500px]">
                        <AdvancedRealTimeChart
                            theme="dark"
                            symbol={symbol}
                            autosize
                            hide_side_toolbar={false}
                            allow_symbol_change={true}
                            studies={[
                                "PivotPointsStandard@tv-basicstudies",
                                "RSI@tv-basicstudies",
                                "MACD@tv-basicstudies"
                            ]}
                        />
                    </div>
                    {!isViewOnly && <StrategyPanel />}
                </div>

                <div className="space-y-6">
                    {!isViewOnly ? (
                        <>
                            <TradePanel exchange={exchange} symbol={symbol} onSymbolChange={setSymbol} />
                            <PositionsPanel />
                            <SignalsPanel />
                        </>
                    ) : (
                        <>
                            <SignalsPanel />
                            <MarketNews />
                        </>
                    )}
                </div>
            </div>
        </div>
    );
};

export default Dashboard;
