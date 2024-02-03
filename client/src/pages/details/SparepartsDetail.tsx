import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Spareparts.module.css';

import { Link } from 'react-router-dom';
import { useParams } from 'react-router-dom';
import { Sparepart } from '../../types/supplierTypes';
import { useUserStore } from '../../lib/store';
import { useGetResourceById } from '../../hooks/genericHooks';

const PartsDetail = async () => {
    const { isAdmin } = useUserStore();

    const id = useParams<{ id: string }>().id as string;
    const spareparts = useGetResourceById<Sparepart[]>('spareparts', id);

    if (!id) {
        return <div>Error: No supplier ID provided</div>;
    }

    return (
        <section id="partsDetail">
            <h1 className={utils.sectionHeading}>Parts Catalogues</h1>

            {spareparts.data && (
                <div className={utils.index}>
                    <h1 className={utils.indexHeading}>spareparts</h1>
                    {spareparts.data.map((link) => (
                        <a key={link.name} href={`#${link.name}`}>
                            <h1 className="indexItem">{link.name}</h1>
                        </a>
                    ))}
                </div>
            )}

            {spareparts.data?.map((sparepart) => (
                <div className={styles.sparepartGrid} key={sparepart.id}>
                    <div className={styles.sparepartsCard} id={sparepart.name || ''}>
                        <div className={styles.sparepartsGrid}>
                            <div className={styles.sparepartsInfo}>
                                <h1 className={utils.mainHeading}>{sparepart.name}</h1>
                                <button className={utils.btn}>
                                    <Link to={sparepart.spare_parts_link || '#'} target="_blank">
                                        Parts Catalogue
                                        <img src="/icons/right-bracket.svg" alt="bracket-right" />
                                    </Link>
                                </button>
                            </div>
                            <img
                                src={sparepart.parts_image || '/default.jpg'}
                                alt={'/default.jpg'}
                                className={styles.sparepartsLogo}
                                width={600}
                                height={600}
                            />
                        </div>
                    </div>
                    {/* {isAdmin && <UpdatePartForm sparepart={sparepart} />} */}
                </div>
            ))}
            {/* {isAdmin && <SparepartsForm />} */}
        </section>
    );
};

export default PartsDetail;