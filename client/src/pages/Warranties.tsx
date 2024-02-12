import utils from '../styles/Utils.module.css';
import styles from '../styles/Account.module.css';

import WarrantyForm from '../forms/WarrantyForm';

import { Link } from 'react-router-dom';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';
import { WarrantyClaim } from '../types/miscTypes';
import { Fragment } from 'react';

import Error from '../layouts/Error';
import Loading from '../layouts/Loading';
import LoginForm from '../forms/LoginForm';

const Warranties: React.FC = () => {
    const { isAdmin, isAuthenticated } = useUserStore();
    const { data: warranties, isLoading, isError } = useGetResource<WarrantyClaim[]>('warranty');

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

    return (
        <section id="warranty">
            {isAuthenticated ? (
                <Fragment>
                    <h1 className={utils.sectionHeading}>Warranty Claims:</h1>
                    <WarrantyForm />

                    {isAdmin &&
                        warranties &&
                        warranties.map(warranty => (
                            <div className={styles.warrantyView} key={warranty.id}>
                                <h1 className={utils.mainHeading}>
                                    {warranty.dealer} -- {warranty.owner_name}
                                </h1>
                                <button className={utils.btnForm}>
                                    <Link to={`/warranty/${warranty.id}`}>View Claim</Link>
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

export default Warranties;
