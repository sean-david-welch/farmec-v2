import utils from '../styles/Utils.module.css';
import styles from '../styles/Account.module.css';
import LogoutButton from '../forms/LogoutButton';
import CarouselAdmin from '../components/CarouselAdmin';
import { useUserStore } from '../lib/store';
import { Fragment } from 'react';
import RegistrationForm from '../forms/RegistrationForm';
import WarrantyForm from '../forms/WarrantyForm';

const Account: React.FC = () => {
    const { isAdmin, isAuthenticated } = useUserStore();
    return (
        <div className={styles.accountMap}>
            <h1 className={utils.sectionHeading}>Account</h1>

            <Fragment>
                {isAdmin ? (
                    <div className={styles.accountSection}>
                        <CarouselAdmin isAdmin={isAdmin} />
                    </div>
                ) : isAuthenticated ? (
                    <div className={styles.accountSection}>
                        <WarrantyForm />
                        <RegistrationForm />
                    </div>
                ) : null}
            </Fragment>
            <LogoutButton mode="button" />
        </div>
    );
};

export default Account;
