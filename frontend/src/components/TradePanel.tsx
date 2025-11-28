import React, { useState } from 'react';
import axios from 'axios';
import { Play, TrendingUp } from 'lucide-react';

interface TradePanelProps {
    exchange: string;
    symbol: string;
    onSymbolChange: (symbol: string) => void;
}

const TradePanel: React.FC<TradePanelProps> = ({ exchange, symbol, onSymbolChange }) => {
    // const [symbol, setSymbol] = useState('BTCUSDT'); // Lifted to parent
    const [quantity, setQuantity] = useState(0.01);
    const [prediction, setPrediction] = useState<number | null>(null);

    const executeTrade = async (side: 'BUY' | 'SELL') => {
        // try {
        //     const res = await axios.post('http://localhost:3000/api/trade', {
        //         exchange,
        //         symbol,
        //         side,
        //         quantity,
        //     });
        //     setPrediction(res.data.profit_prediction);
        //     alert(`Order Executed! ID: ${res.data.orderId}`);
        // } catch (err) {
        //     alert('Trade failed');
        // }
        const mockPrediction = (quantity * 50000) * 0.05; // 5% profit mock
        setPrediction(mockPrediction);
        alert(`Order Executed! ID: mock-order-${Date.now()}`);
    };

    return (
        <div className="bg-surface p-6 rounded-lg shadow-lg">
            <h3 className="text-xl font-bold mb-4 flex items-center gap-2">
                <Play size={20} className="text-accent" />
                Trade Execution ({exchange})
            </h3>

            <div className="space-y-4">
                <div>
                    <label className="block text-sm text-secondary mb-1">Symbol</label>
                    <input
                        type="text"
                        value={symbol}
                        onChange={(e) => onSymbolChange(e.target.value)}
                        className="w-full p-2 rounded bg-background border border-secondary"
                    />
                </div>

                <div>
                    <label className="block text-sm text-secondary mb-1">Quantity</label>
                    <input
                        type="number"
                        value={quantity}
                        onChange={(e) => setQuantity(parseFloat(e.target.value))}
                        className="w-full p-2 rounded bg-background border border-secondary"
                    />
                </div>

                <div className="grid grid-cols-2 gap-4 pt-4">
                    <button
                        onClick={() => executeTrade('BUY')}
                        className="p-3 bg-accent rounded hover:bg-emerald-600 font-bold"
                    >
                        BUY
                    </button>
                    <button
                        onClick={() => executeTrade('SELL')}
                        className="p-3 bg-danger rounded hover:bg-red-600 font-bold"
                    >
                        SELL
                    </button>
                </div>

                {prediction !== null && (
                    <div className="mt-4 p-3 bg-background rounded border border-accent/30">
                        <div className="flex items-center gap-2 text-accent">
                            <TrendingUp size={18} />
                            <span className="font-bold">Predicted Profit: ${prediction}</span>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
};

export default TradePanel;
