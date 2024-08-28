import styles from './styles/App.module.css';

import AppRoutes from './routes/Router';

import {QueryClient, QueryClientProvider} from '@tanstack/react-query';
import useFirebaseAuthSync from './hooks/authHooks';
import {Helmet} from "react-helmet";

function App() {
    useFirebaseAuthSync();
    const queryClient = new QueryClient();

    return (
        <>
            <Helmet>
                <script type="application/ld+json">
                    {JSON.stringify({
                        "@context": "https://schema.org",
                        "@type": "Organization",
                        "name": "Farmec Ireland Ltd",
                        "url": "https://www.farmec.ie",
                        "logo": "https://www.farmec.ie/farmec_images/farmeclogo.webp",
                        "sameAs": [
                            "https://www.facebook.com/farmec",
                            "https://www.twitter.com/farmec1"
                        ],
                        "contactPoint": {
                            "@type": "ContactPoint",
                            "telephone": "+353-1-8259289",
                            "contactType": "Customer Service"
                        }
                    })}
                </script>
            </Helmet>
            <QueryClientProvider client={queryClient}>
                <div className={styles.App}>
                    <AppRoutes/>
                </div>
            </QueryClientProvider>
        </>
    );
}

export default App;
