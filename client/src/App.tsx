import './App.css';
import { ConnectionManager } from './components/connection-manager';
import { ConnectionState } from './components/connection-state';
import { Scale } from './components/scale';
import { socket } from './socket';
import { useState, useEffect } from 'react';

export default function App() {
  const [isConnected, setIsConnected] = useState(socket.connected);

  useEffect(() => {
    function onConnect() {
      setIsConnected(true);
    }

    function onDisconnect() {
      setIsConnected(false);
    }

    socket.on('connect', onConnect);
    socket.on('disconnect', onDisconnect);

    return () => {
      socket.off('connect', onConnect);
      socket.off('disconnect', onDisconnect);
    };
  }, []);

  return (
    <div className="App">
      <Scale />
      <ConnectionState isConnected={isConnected} />
      <ConnectionManager />
    </div>
  );
}
