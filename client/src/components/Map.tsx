import styles from '../styles/Home.module.css';
import config from '../lib/env';

import { GoogleMap, LoadScript, Marker } from '@react-google-maps/api';

export const Map: React.FC = () => {
    const mapsKey = config.mapsKey;

    const containerStyle = {
        width: '600px',
        height: '600px',
    };

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
                <GoogleMap mapContainerStyle={containerStyle} center={center} zoom={10}>
                    <Marker position={location} />
                </GoogleMap>
            </LoadScript>
        </div>
    );
};
