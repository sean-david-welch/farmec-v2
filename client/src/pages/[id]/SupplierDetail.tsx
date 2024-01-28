import utils from '../styles/Utils.module.css';
import styles from '../styles/Suppliers.module.css';

import Videos from '../../templates/Videos';
import Machines from '../../templates/Machines';

import { useParams } from 'react-router-dom';

import { useSupplierDetails } from '../../hooks/supplierHooks';

const SuppliersDetails: React.FC = () => {
    const params = useParams<{ id: string }>();

    if (!params.id) {
        return <div>Error: No supplier ID provided</div>;
    }

    const { supplier, machines, videos } = useSupplierDetails(params.id);

    if (supplier.isLoading || machines.isLoading || videos.isLoading) {
        return <div>Loading...</div>;
    }

    if (supplier.error || machines.error || videos.error) {
        return <div>Error loading data</div>;
    }

    return (
        <section id="supplierDetail">
            {supplier.data && (
                <>
                    <div className={styles.supplierHeading}>
                        <h1 className={utils.sectionHeading}>{supplier.data.name}</h1>
                    </div>

                    {machines.data && (
                        <div className={utils.index}>
                            <h1 className={utils.indexHeading}>Suppliers</h1>
                            {machines.data.map(link => (
                                <a key={link.name} href={`#${link.name}`}>
                                    <h1 className="indexItem">{link.name}</h1>
                                </a>
                            ))}
                        </div>
                    )}

                    <div className={styles.supplierDetail}>
                        <img
                            src={supplier.data.marketing_image ?? '/default.jpg'}
                            alt={'/dafault.jpg'}
                            className={styles.supplierImage}
                            width={750}
                            height={750}
                        />

                        <p className={styles.supplierDescription}>{supplier.data.description}</p>
                    </div>
                </>
            )}

            {machines.data && <Machines machines={machines.data} />}
            {videos.data && <Videos videos={videos.data} />}
        </section>
    );
};

export default SuppliersDetails;
