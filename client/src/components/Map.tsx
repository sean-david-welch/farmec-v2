import styles from '../styles/Home.module.css';
import config from '../lib/env';

import { useState, useEffect } from 'react';
import { GoogleMap, LoadScript, Marker } from '@react-google-maps/api';

export const Map: React.FC = () => {
    const mapsKey = config.mapsKey;

    const [containerStyle, setContainerStyle] = useState({
        width: '600px',
        height: '600px',
    });

    useEffect(() => {
        function handleResize() {
            if (window.innerWidth < 768) {
                setContainerStyle({
                    width: '100%',
                    height: '400px',
                });
            } else {
                setContainerStyle({
                    width: '600px',
                    height: '600px',
                });
            }
        }

        window.addEventListener('resize', handleResize);
        handleResize();

        return () => window.removeEventListener('resize', handleResize);
    }, []);

    const location = {
        lat: 53.49200990196934,
        lng: -6.5423895598058435,
    };

    const center = {
        lat: 53.49200990196934,
        lng: -6.5423895598058435,
    };
    return (
        <div className={styles.map}>
            <LoadScript googleMapsApiKey={mapsKey}>
                <GoogleMap mapContainerStyle={containerStyle} center={center} zoom={12}>
                    <Marker position={location} />
                </GoogleMap>
            </LoadScript>
        </div>
    );
};
