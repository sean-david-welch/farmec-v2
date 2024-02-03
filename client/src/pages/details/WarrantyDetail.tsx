import styles from '../../styles/Services.module.css';
import utils from '../../styles/Utils.module.css';

import { useParams } from 'react-router-dom';
import { DownloadLink } from '../../components/WarrantyPdf';
import { useUserStore } from '../../lib/store';
import { WarrantyParts } from '../../types/miscTypes';
import { useGetResourceById } from '../../hooks/genericHooks';

const WarrantyDetail = async () => {
    const { isAdmin } = useUserStore();

    const id = useParams<{ id: string }>().id as string;
    const response = useGetResourceById<WarrantyParts>('warranty', id);

    if (!id) {
        return <div>Error: No supplier ID provided</div>;
    }

    const parts = response.data?.parts;
    const warranty = response.data?.warranty;

    if (!warranty || !parts) {
        return (
            <section id="warranty-detail">
                <div>Warranty claim not found</div>
            </section>
        );
    }

    return (
        <section id="warranty-detail">
            <h1 className={utils.sectionHeading}>
                Warranty Claim: {warranty.dealer} - {warranty.owner_name}
            </h1>

            <div className={styles.warrantyDetail}>
                {Object.entries(warranty).map(([key, value]) => {
                    if (key !== 'id' && key !== 'created' && key !== 'parts') {
                        return (
                            <div className={styles.warrantyGrid} key={key}>
                                <div className={styles.label}>{key}</div>
                                <div className={styles.value}>{String(value)}</div>
                            </div>
                        );
                    }
                })}
                {parts.map((part, index) =>
                    Object.entries(part).map(([key, value]) => {
                        if (key !== 'id' && key !== 'warrantyId') {
                            return (
                                <div className={styles.warrantyGrid} key={key + index}>
                                    <div className={styles.label}>{key}</div>
                                    <div className={styles.value}>{String(value)}</div>
                                </div>
                            );
                        }
                    })
                )}

                {/* {isAdmin && <UpdateWarranty warrantyClaim={warranty} partsRequired={parts} />} */}
            </div>

            <DownloadLink warranty={warranty} parts={parts} />
        </section>
    );
};

export default WarrantyDetail;
