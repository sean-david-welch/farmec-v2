import styles from '../../styles/Account.module.css';
import utils from '../../styles/Utils.module.css';

import Error from '../../layouts/Error';
import Loading from '../../layouts/Loading';
import WarrantyForm from '../../forms/WarrantyForm';
import DeleteButton from '../../components/DeleteButton';

import { useParams } from 'react-router-dom';
import { DownloadLink } from '../../components/WarrantyPdf';
import { useUserStore } from '../../lib/store';
import { WarrantyParts } from '../../types/miscTypes';
import { useGetResourceById } from '../../hooks/genericHooks';

const WarrantyDetail: React.FC = () => {
    const { isAdmin } = useUserStore();

    const id = useParams<{ id: string }>().id as string;
    const { data, isError, isLoading } = useGetResourceById<WarrantyParts>('warranty', id);

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

    if (!id) {
        return <div>Error: No supplier ID provided</div>;
    }

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

                {isAdmin && warranty.id && (
                    <div className={utils.optionsBtn}>
                        <WarrantyForm id={warranty.id} warranty={warranty} />
                        <DeleteButton id={warranty.id} resourceKey="warranty" navigateBack={true} />
                    </div>
                )}
            </div>

            <DownloadLink warranty={warranty} parts={parts} />
        </section>
    );
};

export default WarrantyDetail;
