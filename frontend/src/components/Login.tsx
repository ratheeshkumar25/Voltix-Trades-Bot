import React, { useState } from 'react';
import { GoogleLogin, GoogleOAuthProvider } from '@react-oauth/google';
import axios from 'axios';
import logo from '../assets/voltix-logo.jpg';

interface LoginProps {
    onLogin: (token: string) => void;
}

const API_URL = import.meta.env.VITE_AUTH_API_URL || 'http://localhost:3001/api';
const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID || '';

type LoginMode = 'select' | 'email' | 'register';

const Login: React.FC<LoginProps> = ({ onLogin }) => {
    const [mode, setMode] = useState<LoginMode>('select');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);

    const handleEmailLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setLoading(true);

        try {
            const response = await axios.post(`${API_URL}/auth/login`, {
                email,
                password
            });
            onLogin(response.data.token);
        } catch (err: any) {
            setError(err.response?.data?.error || 'Login failed');
        } finally {
            setLoading(false);
        }
    };

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setLoading(true);

        try {
            const response = await axios.post(`${API_URL}/auth/register`, {
                email,
                password
            });
            onLogin(response.data.token);
        } catch (err: any) {
            setError(err.response?.data?.error || 'Registration failed');
        } finally {
            setLoading(false);
        }
    };

    const handleGoogleSuccess = async (credentialResponse: any) => {
        setError('');
        setLoading(true);

        try {
            // Send Google credential to backend
            const response = await axios.post(`${API_URL}/auth/google`, {
                credential: credentialResponse.credential
            });
            onLogin(response.data.token);
        } catch (err: any) {
            setError(err.response?.data?.error || 'Google login failed');
        } finally {
            setLoading(false);
        }
    };

    return (
        <GoogleOAuthProvider clientId={GOOGLE_CLIENT_ID}>
            <div className="flex items-center justify-center min-h-screen bg-gray-900 p-4">
                <div className="absolute inset-0 bg-gradient-to-br from-blue-900/20 to-purple-900/20" />

                {mode === 'select' ? (
                    // Mode Selection
                    <div className="relative z-10 w-full max-w-md">
                        <div className="text-center mb-8">
                            <img src={logo} alt="Voltix Logo" className="h-20 w-auto mx-auto mb-4 rounded-lg" />
                            <h1 className="text-4xl font-bold text-white mb-2">Welcome to Voltix</h1>
                            <p className="text-gray-300">Start your 7-day free trial</p>
                        </div>

                        <div className="bg-surface/80 backdrop-blur-md rounded-xl p-8 border border-white/10">
                            {/* Google Sign In */}
                            <div className="mb-6">
                                <GoogleLogin
                                    onSuccess={handleGoogleSuccess}
                                    onError={() => setError('Google login failed')}
                                    theme="filled_black"
                                    size="large"
                                    width="100%"
                                />
                            </div>

                            <div className="relative my-6">
                                <div className="absolute inset-0 flex items-center">
                                    <div className="w-full border-t border-gray-600"></div>
                                </div>
                                <div className="relative flex justify-center text-sm">
                                    <span className="px-4 bg-surface text-gray-400">Or continue with</span>
                                </div>
                            </div>

                            {/* Email/Password Options */}
                            <div className="space-y-3">
                                <button
                                    onClick={() => setMode('email')}
                                    className="w-full p-3 bg-primary hover:bg-blue-600 rounded-lg font-bold transition-colors"
                                >
                                    Sign in with Email
                                </button>
                                <button
                                    onClick={() => setMode('register')}
                                    className="w-full p-3 bg-background hover:bg-surface border border-secondary rounded-lg font-bold transition-colors"
                                >
                                    Create Account
                                </button>
                            </div>

                            {error && (
                                <div className="mt-4 p-3 bg-red-500/20 border border-red-500 rounded text-red-500 text-sm">
                                    {error}
                                </div>
                            )}
                        </div>
                    </div>
                ) : (
                    // Email Login/Register Form
                    <div className="relative z-10 w-full max-w-md">
                        <div className="bg-surface/80 backdrop-blur-md rounded-xl p-8 border border-white/10">
                            <button
                                onClick={() => {
                                    setMode('select');
                                    setError('');
                                }}
                                className="text-sm text-gray-400 hover:text-white mb-6 flex items-center gap-1"
                            >
                                ← Back
                            </button>

                            <div className="text-center mb-6">
                                <h2 className="text-2xl font-bold text-white mb-2">
                                    {mode === 'email' ? 'Sign In' : 'Create Account'}
                                </h2>
                                <p className="text-sm text-gray-400">
                                    {mode === 'email'
                                        ? 'Enter your credentials to continue'
                                        : 'Start your 7-day free trial'}
                                </p>
                            </div>

                            {error && (
                                <div className="mb-4 p-3 bg-red-500/20 border border-red-500 rounded text-red-500 text-sm">
                                    {error}
                                </div>
                            )}

                            <form onSubmit={mode === 'email' ? handleEmailLogin : handleRegister} className="space-y-4">
                                <div>
                                    <label className="block text-sm font-medium mb-2 text-gray-300">Email</label>
                                    <input
                                        type="email"
                                        value={email}
                                        onChange={(e) => setEmail(e.target.value)}
                                        placeholder="your.email@example.com"
                                        className="w-full p-3 rounded-lg bg-background border border-secondary focus:border-primary outline-none transition-colors"
                                        required
                                        disabled={loading}
                                    />
                                </div>

                                <div>
                                    <label className="block text-sm font-medium mb-2 text-gray-300">Password</label>
                                    <input
                                        type="password"
                                        value={password}
                                        onChange={(e) => setPassword(e.target.value)}
                                        placeholder="••••••••"
                                        className="w-full p-3 rounded-lg bg-background border border-secondary focus:border-primary outline-none transition-colors"
                                        required
                                        disabled={loading}
                                        minLength={6}
                                    />
                                </div>

                                <button
                                    type="submit"
                                    disabled={loading}
                                    className="w-full p-3 bg-primary hover:bg-blue-600 rounded-lg font-bold transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                                >
                                    {loading ? 'Please wait...' : (mode === 'email' ? 'Sign In' : 'Create Account')}
                                </button>
                            </form>

                            <div className="mt-6 text-center">
                                <button
                                    onClick={() => setMode(mode === 'email' ? 'register' : 'email')}
                                    className="text-sm text-gray-400 hover:text-white"
                                >
                                    {mode === 'email'
                                        ? "Don't have an account? Sign up"
                                        : 'Already have an account? Sign in'}
                                </button>
                            </div>
                        </div>
                    </div>
                )}
            </div>
        </GoogleOAuthProvider>
    );
};

export default Login;
