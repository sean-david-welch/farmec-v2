import utils from '../styles/Utils.module.css';
import styles from '../styles/Account.module.css';

import WarrantyForm from '../forms/WarrantyForm';

import { Link } from 'react-router-dom';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';
import { WarrantyClaim } from '../types/miscTypes';

const Warranties: React.FC = () => {
    const { isAdmin } = useUserStore();
    const { data: warranties, isLoading } = useGetResource<WarrantyClaim[]>('warranty');

    if (isLoading) return <div>Loeading...</div>;

    return (
        <section id="warranty">
            <h1 className={utils.sectionHeading}>Warranty Claims:</h1>
            {warranties &&
                warranties.map((warranty) => (
                    <div className={styles.warrantyView} key={warranty.id}>
                        <h1 className={utils.mainHeading}>
                            {warranty.dealer} -- {warranty.owner_name}
                        </h1>
                        <button className={utils.btnForm}>
                            <Link to={`/warranty/${warranty.id}`}>View Claim</Link>
                        </button>
                    </div>
                ))}
            {isAdmin && <WarrantyForm />}
        </section>
    );
};

export default Warranties;
