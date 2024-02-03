import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Suppliers.module.css';

import Videos from '../../templates/Videos';
import Machines from '../../templates/Machines';

import { useParams } from 'react-router-dom';

import { useMultipleResources } from '../../hooks/genericHooks';
import { Resources } from '../../types/dataTypes';
import { useUserStore } from '../../lib/store';

const SuppliersDetails: React.FC = () => {
    const { isAdmin } = useUserStore();

    const id = useParams<{ id: string }>().id as string;

    const resourceKeys: (keyof Resources)[] = ['suppliers', 'supplierMachine', 'videos'];
    const { data, isLoading } = useMultipleResources(id, resourceKeys);

    if (!id) {
        return <div>Error: No supplier ID provided</div>;
    }

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

            {machines.length > 0 && <Machines machines={machines} isAdmin={isAdmin} />}
            {videos.length > 0 && <Videos videos={videos} />}
        </section>
    );
};

export default SuppliersDetails;
