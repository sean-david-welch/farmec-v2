import utils from '../styles/Utils.module.css';
import styles from '../styles/Account.module.css';

import Error from '../layouts/Error';
import Loading from '../layouts/Loading';
import LoginForm from '../forms/LoginForm';
import RegistrationForm from '../forms/RegistrationForm';

import { Link } from 'react-router-dom';
import { Fragment } from 'react';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';
import { MachineRegistration } from '../types/miscTypes';

const Registrations: React.FC = () => {
    const { isAdmin, isAuthenticated } = useUserStore();
    const { data: registrations, isLoading, isError } = useGetResource<MachineRegistration[]>('registrations');

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

    return (
        <section id="registrations">
            {isAuthenticated ? (
                <Fragment>
                    <h1 className={utils.sectionHeading}>Machine Registration:</h1>
                    <RegistrationForm />
                    {isAdmin &&
                        registrations &&
                        registrations.map(registration => (
                            <div className={styles.warrantyView} key={registration.id}>
                                <h1 className={utils.mainHeading}>
                                    {registration.dealer_name} -- {registration.owner_name}
                                </h1>
                                <button className={utils.btnForm}>
                                    <Link to={`/registration/${registration.id}`}>View Claim</Link>
                                </button>
                            </div>
                        ))}
                </Fragment>
            ) : (
                <div className={utils.loginSection}>
                    <LoginForm />
                </div>
            )}
        </section>
    );
};

export default Registrations;
