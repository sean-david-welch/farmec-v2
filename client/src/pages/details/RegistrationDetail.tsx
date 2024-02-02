import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Account.module.css';

import { useParams } from 'react-router-dom';
import { useUserStore } from '../../lib/store';
import { useGetResourceById } from '../../hooks/genericHooks';
import { MachineRegistration } from '../../types/miscTypes';
import { DownloadLink } from '../../components/RegistrationPdf';

const RegistrationDetail: React.FC = () => {
    const isAdmin = useUserStore();

    const params = useParams<{ id: string }>();
    const id = params.id as string;
    const registration = useGetResourceById<MachineRegistration>('registrations', id);

    if (!params.id) {
        return <div>Error: No supplier ID provided</div>;
    }

    if (!registration) {
        return (
            <section id="warranty-detail">
                <div>Warranty claim not found</div>
            </section>
        );
    }

    return (
        registration.data && (
            <section id="warranty-detail">
                <h1 className={utils.sectionHeading}>
                    Machine Registration: {registration.data?.dealer_name} - {registration.data?.owner_name}
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

                    {/* {isAdmin && <UpdateRegistration registration={registration} />} */}
                </div>

                <DownloadLink registration={registration.data} />
            </section>
        )
    );
};

export default RegistrationDetail;
