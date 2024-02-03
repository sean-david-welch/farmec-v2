import styles from '../styles/About.module.css';
import utils from '../styles/Utils.module.css';

import { useGetResource } from '../hooks/genericHooks';
import { Employee } from '../types/aboutTypes';
import { useUserStore } from '../lib/store';

export const Employees = () => {
    const { isAdmin } = useUserStore();

    const employees = useGetResource<Employee[]>('employees');

    return (
        <section id="employees">
            <div className={styles.employees}>
                {employees.data?.map((employee) => (
                    <div className={styles.employeeCard} key={employee.id}>
                        <img src={employee.profile_image || '/default.jpg'} alt={'/default.jpg'} width={100} height={100} />
                        <div className={styles.employeeInfo}>
                            <h1 className={utils.mainHeading}>{employee.name}</h1>
                            <p className={utils.paragraph}>{employee.role}</p>
                        </div>
                        <div className={styles.employeeContact}>
                            <p className={utils.paragraph}>Email: {employee.email}</p>
                        </div>
                        {/* {isAdmin && <UpdateAbout modelName="employee" model={employee} />} */}
                    </div>
                ))}
            </div>
        </section>
    );
};
