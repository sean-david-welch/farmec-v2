import styles from './styles/App.module.css';

import AppRoutes from './routes/Router';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

function App() {
    const queryClient = new QueryClient();

    return (
        <QueryClientProvider client={queryClient}>
            <div className={styles.App}>
                <AppRoutes />
            </div>
        </QueryClientProvider>
    );
}

export default App;
