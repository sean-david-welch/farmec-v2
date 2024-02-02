import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Suppliers.module.css';

import Videos from '../../templates/Videos';
import Machines from '../../templates/Machines';

import { useParams } from 'react-router-dom';

import { useMultipleResources } from '../../hooks/genericHooks';
import { Resources } from '../../types/dataTypes';

const SuppliersDetails: React.FC = () => {
    const params = useParams<{ id: string }>();

    if (!params.id) {
        return <div>Error: No supplier ID provided</div>;
    }

    const resourceKeys: (keyof Resources)[] = ['suppliers', 'machines', 'videos'];
    const { data, isLoading } = useMultipleResources(params.id, resourceKeys);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    const [supplier, machines, videos] = data;

    return (
        <section id="supplierDetail">
            {supplier && (
                <>
                    <div className={styles.supplierHeading}>
                        <h1 className={utils.sectionHeading}>{supplier.name}</h1>
                    </div>

                    {machines.length > 0 && (
                        <div className={utils.index}>
                            <h1 className={utils.indexHeading}>Suppliers</h1>
                            {machines.map((link: { name: string }) => (
                                <a key={link.name} href={`#${link.name}`}>
                                    <h1 className="indexItem">{link.name}</h1>
                                </a>
                            ))}
                        </div>
                    )}

                    <div className={styles.supplierDetail}>
                        <img
                            src={supplier.marketing_image ?? '/default.jpg'}
                            alt={'/dafault.jpg'}
                            className={styles.supplierImage}
                            width={750}
                            height={750}
                        />

                        <p className={styles.supplierDescription}>{supplier.description}</p>
                    </div>
                </>
            )}

            {machines.length > 0 && <Machines machines={machines} />}
            {videos.length > 0 && <Videos videos={videos} />}
        </section>
    );
};

export default SuppliersDetails;
