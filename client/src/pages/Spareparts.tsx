import utils from '../styles/Utils.module.css';
import styles from '../styles/Suppliers.module.css';

import { Link } from 'react-router-dom';

import { useSupplierStore, useUserStore } from '../lib/store';
import SparepartForm from '../forms/SparePartsForm';

const SpareParts: React.FC = () => {
    const { isAdmin } = useUserStore();

    const { suppliers } = useSupplierStore();

    return (
        <section id="SpareParts">
            <h1 className={utils.sectionHeading}>Spare-Parts</h1>

            {suppliers && (
                <div className={utils.index}>
                    <h1 className={utils.indexHeading}>Suppliers</h1>
                    {suppliers.map((link) => (
                        <a key={link.name} href={`#${link.name}`}>
                            <h1 className="indexItem">{link.name}</h1>
                        </a>
                    ))}
                </div>
            )}
            {suppliers.map((supplier) => (
                <div className={styles.supplierCard} key={supplier.id}>
                    <h1 className={utils.mainHeading}>{supplier.name}</h1>
                    <img
                        src={supplier.logo_image ?? '/default.jpg'}
                        alt="Supplier logo"
                        width={200}
                        height={200}
                    />
                    <button className={utils.btn}>
                        <Link to={`/spareparts/${supplier.id}`}>
                            Spare-Parts
                            <img src="/icons/right-bracket.svg" alt="bracket-right" />
                        </Link>
                    </button>
                </div>
            ))}
            {isAdmin && suppliers && <SparepartForm suppliers={suppliers} />}
        </section>
    );
};

export default SpareParts;
