import styles from '../../styles/Account.module.css';
import utils from '../../styles/Utils.module.css';

import ErrorPage from '../../layouts/Error';
import Loading from '../../layouts/Loading';
import WarrantyForm from '../../forms/WarrantyForm';
import DeleteButton from '../../components/DeleteButton';
import DownloadPdfButton from '../../components/DownloadPdf';

import { useEffect } from 'react';
import { useParams } from 'react-router-dom';

import { useUserStore } from '../../lib/store';
import { WarrantyParts } from '../../types/miscTypes';
import { useGetResourceById } from '../../hooks/genericHooks';

const WarrantyDetail: React.FC = () => {
    const { isAdmin } = useUserStore();

    const id = useParams<{ id: string }>().id as string;
    const { data, isError, isLoading } = useGetResourceById<WarrantyParts>('warranty', id);

    useEffect(() => {}, [id]);

    if (isError) return <ErrorPage />;
    if (isLoading) return <Loading />;

    const parts = data?.parts;
    const warranty = data?.warranty;

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
                        if (key !== 'id' && key !== 'warranty_id') {
                            return (
                                <div className={styles.warrantyGrid} key={key + index}>
                                    <div className={styles.label}>{key}</div>
                                    <div className={styles.value}>{String(value)}</div>
                                </div>
                            );
                        }
                    })
                )}

                {isAdmin && warranty.id && (
                    <div className={utils.optionsBtn}>
                        <WarrantyForm id={warranty.id} warranty={warranty} />
                        <DeleteButton id={warranty.id} resourceKey="warranty" navigateBack={true} />
                    </div>
                )}
            </div>

            <DownloadPdfButton warrantyClaim={warranty} partsRequired={parts} />
        </section>
    );
};

export default WarrantyDetail;
