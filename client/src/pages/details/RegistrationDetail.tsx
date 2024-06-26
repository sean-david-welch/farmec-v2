import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Account.module.css';

import ErrorPage from '../../layouts/Error';
import Loading from '../../layouts/Loading';
import RegistrationForm from '../../forms/RegistrationForm';
import DeleteButton from '../../components/DeleteButton';

import { useParams } from 'react-router-dom';
import { useUserStore } from '../../lib/store';
import { useGetResourceById } from '../../hooks/genericHooks';
import { MachineRegistration } from '../../types/miscTypes';

import { useEffect } from 'react';
import DownloadPdfButton from '../../components/DownloadPdf';

const RegistrationDetail: React.FC = () => {
    const { isAdmin } = useUserStore();

    const id = useParams<{ id: string }>().id as string;
    const { data: registration, isLoading, isError } = useGetResourceById<MachineRegistration>('registrations', id);

    useEffect(() => {}, [id]);

    if (isError) return <ErrorPage />;
    if (isLoading) return <Loading />;

    if (!registration) {
        return (
            <section id="warranty-detail">
                <div>Warranty claim not found</div>
            </section>
        );
    }

    return (
        registration && (
            <section id="warranty-detail">
                <h1 className={utils.sectionHeading}>
                    Machine Registration: {registration?.dealer_name} - {registration?.owner_name}
                </h1>

                <div className={styles.warrantyDetail}>
                    {Object.entries(registration).map(([key, value]) => {
                        if (key !== 'id' && key !== 'created' && key !== 'parts') {
                            return (
                                <div className={styles.warrantyGrid} key={key}>
                                    <div className={styles.label}>{key}</div>
                                    <div className={styles.value}>{String(value)}</div>
                                </div>
                            );
                        }
                    })}

                    {isAdmin && registration.id && (
                        <div className={utils.optionsBtn}>
                            <RegistrationForm id={registration.id} />
                            <DeleteButton id={registration.id} resourceKey="registrations" navigateBack={true} />
                        </div>
                    )}
                </div>

                <DownloadPdfButton registration={registration} />
            </section>
        )
    );
};

export default RegistrationDetail;
