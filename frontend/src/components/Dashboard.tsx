import React, { useState, useEffect } from 'react';
import axios from 'axios';
// import { LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer } from 'recharts';
import { AdvancedRealTimeChart } from "react-ts-tradingview-widgets";
import { LayoutDashboard, Wallet, ArrowRightLeft, Coins } from 'lucide-react';
import TradePanel from './TradePanel';
import StrategyPanel from './StrategyPanel';
import PositionsPanel from './PositionsPanel';
import SignalsPanel from './SignalsPanel';
import logo from '../assets/voltix-logo.jpg';

const Dashboard: React.FC = () => {
    const [exchange, setExchange] = useState('binance');
    const [symbol, setSymbol] = useState('BTCUSDT');
    const [balance, setBalance] = useState(0);
    const [data] = useState([
        { name: '10:00', price: 50000 },
        { name: '10:05', price: 50200 },
        { name: '10:10', price: 50100 },
        { name: '10:15', price: 50400 },
        { name: '10:20', price: 50300 },
        { name: '10:25', price: 50600 },
    ]);

    useEffect(() => {
        const fetchBalance = async () => {
            // try {
            //     const res = await axios.get(`http://localhost:3000/api/balance/${exchange}`);
            //     setBalance(res.data.balance);
            // } catch (err) {
            //     console.error(err);
            // }
            setBalance(10000); // Static balance
        };
        fetchBalance();
    }, [exchange]);

    return (
        <div className="min-h-screen bg-background p-6">
            <header className="flex justify-between items-center mb-8">
                <h1 className="text-3xl font-bold text-primary flex items-center gap-2">
                    <img src={logo} alt="Voltix Logo" className="h-10 w-auto rounded" />
                    Voltix Dashboard
                </h1>
                <div className="flex items-center gap-4">
                    <div className="flex items-center gap-2 bg-surface px-4 py-2 rounded-lg">
                        <Wallet size={18} className="text-secondary" />
                        <span className="font-mono text-lg">${balance.toFixed(2)}</span>
                    </div>
                    <div className="flex items-center gap-2 bg-surface px-4 py-2 rounded-lg">
                        <Coins size={18} className="text-secondary" />
                        <select
                            value={symbol}
                            onChange={(e) => setSymbol(e.target.value)}
                            className="bg-transparent outline-none font-bold text-accent"
                        >
                            <option value="BTCUSDT">BTC/USDT</option>
                            <option value="ETHUSDT">ETH/USDT</option>
                            <option value="SOLUSDT">SOL/USDT</option>
                            <option value="XRPUSDT">XRP/USDT</option>
                        </select>
                    </div>
                    <div className="flex items-center gap-2 bg-surface px-4 py-2 rounded-lg">
                        <ArrowRightLeft size={18} className="text-secondary" />
                        <select
                            value={exchange}
                            onChange={(e) => setExchange(e.target.value)}
                            className="bg-transparent outline-none font-bold text-accent"
                        >
                            <option value="binance">Binance</option>
                            <option value="mt5">MetaTrader 5</option>
                            <option value="ctrader">cTrader</option>
                        </select>
                    </div>
                </div>
            </header>

            <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                <div className="lg:col-span-2 bg-surface p-6 rounded-lg shadow-lg">
                    <h3 className="text-xl font-bold mb-4">Market Overview ({exchange.toUpperCase()})</h3>
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
                    <StrategyPanel />
                </div>

                <div className="space-y-6">
                    <TradePanel exchange={exchange} symbol={symbol} onSymbolChange={setSymbol} />

                    <PositionsPanel />
                    <SignalsPanel />
                </div>
            </div>
        </div>
    );
};

export default Dashboard;
