import utils from '../styles/Utils.module.css';
import styles from '../styles/Account.module.css';

import { User } from '../types/dataTypes';
import { useGetResource } from '../hooks/genericHooks';
import { useUserStore } from '../lib/store';
import { Fragment } from 'react';
import RegisterForm from '../forms/RegisterForm';
import DeleteButton from '../components/DeleteButton';
import Loading from '../layouts/Loading';

const Users: React.FC = () => {
    const { isAdmin } = useUserStore();
    const { data: users, isLoading } = useGetResource<User[]>('users');

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

    return (
        <section id="users">
            <div className={styles.usersSection}>
                <h1 className={utils.sectionHeading}>Users</h1>
                {isAdmin &&
                    users?.map(user => (
                        <Fragment key={user.rawId}>
                            <div className={styles.productView}>
                                <h1 className={utils.paragraph}>
                                    {user.email} -- {user.CustomClaims?.admin ? 'Admin' : 'Not Admin'}
                                </h1>
                                <div className={utils.optionsBtn}>
                                    <RegisterForm id={user.rawId} />
                                    <DeleteButton id={user?.rawId} resourceKey="users" />
                                </div>
                            </div>
                        </Fragment>
                    ))}
                <RegisterForm />
            </div>
        </section>
    );
};

export default Users;
