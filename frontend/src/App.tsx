import { useState } from 'react';
import Login from './components/Login';
import Dashboard from './components/Dashboard';

type AccountType = 'metatrader' | 'binance' | 'ctrader' | 'gmail';

function App() {
  const [token, setToken] = useState<string | null>(null);
  const [accountType, setAccountType] = useState<AccountType | null>(null);

  const handleLogin = (loginToken: string, loginAccountType: AccountType) => {
    setToken(loginToken);
    setAccountType(loginAccountType);
  };

  if (!token || !accountType) {
    return <Login onLogin={handleLogin} />;
  }

  return <Dashboard accountType={accountType} />;
}

export default App;
