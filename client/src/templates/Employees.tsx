import styles from '../styles/About.module.css';
import utils from '../styles/Utils.module.css';

import EmployeeForm from '../forms/EmployeeForm';
import DeleteButton from '../components/DeleteButton';

import { Employee } from '../types/aboutTypes';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';

export const Employees = () => {
	const { isAdmin } = useUserStore();

	const employees = useGetResource<Employee[]>('employees');
	const imageError = (event: React.SyntheticEvent<HTMLImageElement, Event>) => {
		event.currentTarget.src = '/default.jpg';
	};

	return (
		<section id="employees">
			<div className={styles.employees}>
				{employees.data?.map(employee => (
					<div className={styles.employeeCard} key={employee.id}>
						<img
							src={employee.profile_image}
							alt={'/default.jpg'}
							width={100}
							height={100}
							onError={imageError}
						/>
						<div className={styles.employeeInfo}>
							<h1 className={utils.mainHeading}>{employee.name}</h1>
							<p className={utils.paragraph}>{employee.role}</p>
						</div>
						<div className={styles.employeeContact}>
							<p className={utils.paragraph}>Email: {employee.email}</p>
						</div>
						{isAdmin && employee.id && (
							<div className={utils.optionsBtn}>
								<EmployeeForm id={employee.id} employee={employee} />
								<DeleteButton id={employee.id} resourceKey="employees" />
							</div>
						)}
					</div>
				))}
			</div>

			{isAdmin && <EmployeeForm />}
		</section>
	);
};
