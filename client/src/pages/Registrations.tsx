import utils from '../styles/Utils.module.css';
import styles from '../styles/Account.module.css';

import WarrantyForm from '../forms/WarrantyForm';

import { Link } from 'react-router-dom';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';
import { MachineRegistration } from '../types/miscTypes';

const Registrations: React.FC = () => {
    const { isAdmin } = useUserStore();
    const { data: registrations, isLoading } = useGetResource<MachineRegistration[]>('registrations');

    if (isLoading) return <div>Loeading...</div>;

    return (
        <section id="registrations">
            <h1 className={utils.sectionHeading}>Machine Registration:</h1>
            {registrations &&
                registrations.map((registration) => (
                    <div className={styles.warrantyView} key={registration.id}>
                        <h1 className={utils.mainHeading}>
                            {registration.dealer_name} -- {registration.owner_name}
                        </h1>
                        <button className={utils.btnForm}>
                            <Link to={`/registration/${registration.id}`}>View Claim</Link>
                        </button>
                    </div>
                ))}
            {isAdmin && <WarrantyForm />}
        </section>
    );
};

export default Registrations;
