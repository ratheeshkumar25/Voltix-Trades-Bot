import React, { useState } from 'react';
import axios from 'axios';
import logo from '../assets/voltix-logo.jpg';
import bgImage from '../assets/login-bg.jpg';

interface LoginProps {
    onLogin: (token: string, accountType: AccountType) => void;
}

type AccountType = 'metatrader' | 'binance' | 'ctrader' | 'gmail';

const Login: React.FC<LoginProps> = ({ onLogin }) => {
    const [accountType, setAccountType] = useState<AccountType | null>(null);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        // Mock authentication - pass account type to parent
        if (accountType) {
            onLogin(`mock-token-${accountType}`, accountType);
        }
    };

    const accountOptions = [
        {
            id: 'metatrader' as AccountType,
            name: 'MetaTrader',
            description: 'MT5 Trading Account',
            icon: '📈',
            color: 'from-blue-500 to-blue-600'
        },
        {
            id: 'binance' as AccountType,
            name: 'Binance',
            description: 'Exchange Trading',
            icon: '🔶',
            color: 'from-yellow-500 to-yellow-600'
        },
        {
            id: 'ctrader' as AccountType,
            name: 'cTrader',
            description: 'Platform Trading',
            icon: '📊',
            color: 'from-green-500 to-green-600'
        },
        {
            id: 'gmail' as AccountType,
            name: 'Gmail',
            description: 'View Only Access',
            icon: '📧',
            color: 'from-red-500 to-red-600'
        }
    ];

    return (
        <div
            className="flex items-center justify-center min-h-screen bg-gray-900 relative p-4"
            style={{
                backgroundImage: `url(${bgImage})`,
                backgroundSize: 'cover',
                backgroundPosition: 'center',
                backgroundRepeat: 'no-repeat'
            }}
        >
            <div className="absolute inset-0 bg-black/40" />

            {!accountType ? (
                // Account Type Selection
                <div className="relative z-10 w-full max-w-4xl">
                    <div className="text-center mb-8">
                        <img src={logo} alt="Voltix Logo" className="h-20 w-auto mx-auto mb-4 rounded-lg" />
                        <h1 className="text-4xl font-bold text-white mb-2">Welcome to Voltix</h1>
                        <p className="text-gray-300">Choose your account type to continue</p>
                    </div>

                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                        {accountOptions.map((option) => (
                            <button
                                key={option.id}
                                onClick={() => setAccountType(option.id)}
                                className="group relative overflow-hidden bg-surface/80 backdrop-blur-md rounded-xl p-6 border border-white/10 hover:border-white/30 transition-all hover:scale-105 hover:shadow-2xl"
                            >
                                <div className={`absolute inset-0 bg-gradient-to-br ${option.color} opacity-0 group-hover:opacity-10 transition-opacity`} />
                                <div className="relative">
                                    <div className="text-5xl mb-3">{option.icon}</div>
                                    <h3 className="text-xl font-bold text-white mb-1">{option.name}</h3>
                                    <p className="text-sm text-gray-400">{option.description}</p>
                                </div>
                            </button>
                        ))}
                    </div>
                </div>
            ) : (
                // Login Form
                <div className="p-8 bg-surface/80 backdrop-blur-md rounded-lg shadow-2xl w-96 relative z-10 border border-white/10">
                    <button
                        onClick={() => setAccountType(null)}
                        className="text-sm text-gray-400 hover:text-white mb-4 flex items-center gap-1"
                    >
                        ← Back to account selection
                    </button>

                    <div className="flex justify-center mb-4">
                        <div className="text-5xl">
                            {accountOptions.find(opt => opt.id === accountType)?.icon}
                        </div>
                    </div>

                    <h2 className="text-2xl font-bold mb-2 text-center text-primary">
                        {accountOptions.find(opt => opt.id === accountType)?.name} Login
                    </h2>

                    <p className="text-sm text-gray-400 text-center mb-6">
                        {accountType === 'gmail'
                            ? 'View predictions and strategies'
                            : 'Full trading access'}
                    </p>

                    {error && <p className="text-danger mb-4 text-center">{error}</p>}

                    <form onSubmit={handleSubmit} className="space-y-4">
                        <div>
                            <label className="block text-sm font-medium mb-1">
                                {accountType === 'gmail' ? 'Email' : 'Username'}
                            </label>
                            <input
                                type={accountType === 'gmail' ? 'email' : 'text'}
                                value={username}
                                onChange={(e) => setUsername(e.target.value)}
                                placeholder={accountType === 'gmail' ? 'your.email@gmail.com' : 'Enter username'}
                                className="w-full p-2 rounded bg-background border border-secondary focus:border-primary outline-none"
                                required
                            />
                        </div>
                        <div>
                            <label className="block text-sm font-medium mb-1">Password</label>
                            <input
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                placeholder="Enter password"
                                className="w-full p-2 rounded bg-background border border-secondary focus:border-primary outline-none"
                                required
                            />
                        </div>
                        <button
                            type="submit"
                            className="w-full p-2 bg-primary rounded hover:bg-blue-600 transition-colors font-bold"
                        >
                            Login
                        </button>
                    </form>
                </div>
            )}
        </div>
    );
};

export default Login;
