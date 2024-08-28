import utils from '../styles/Utils.module.css';
import styles from '../styles/Suppliers.module.css';

import {Link} from 'react-router-dom';
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome';
import {faRightToBracket} from '@fortawesome/free-solid-svg-icons';
import {useSupplierStore, useUserStore} from '../lib/store';

import WarrantyForm from '../forms/WarrantyForm';
import SparepartForm from '../forms/SparePartsForm';
import RegistrationForm from '../forms/RegistrationForm';
import {Helmet} from 'react-helmet';
import {FC} from "react";

const SpareParts: FC = () => {
	const { isAdmin } = useUserStore();
	const { suppliers } = useSupplierStore();

	return (
		<>
			<Helmet>
				<title>Spare Parts - Farmec Ireland</title>
				<meta name="description" content="Browse spare parts from various suppliers and ensure your machines are always operational. Farmec offers a wide range of spare parts for agricultural machinery." />

				<meta property="og:title" content="Spare Parts - Farmec Ireland" />
				<meta property="og:description" content="Browse spare parts from various suppliers and ensure your machines are always operational. Farmec offers a wide range of spare parts for agricultural machinery." />
				<meta property="og:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />
				<meta property="og:url" content="https://www.farmec.ie/spareparts" />
				<meta property="og:type" content="website" />

				<meta name="twitter:card" content="summary_large_image" />
				<meta name="twitter:title" content="Spare Parts - Farmec Ireland" />
				<meta name="twitter:description" content="Browse spare parts from various suppliers and ensure your machines are always operational. Farmec offers a wide range of spare parts for agricultural machinery." />
				<meta name="twitter:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />

				<link rel="canonical" href="https://www.farmec.ie/spareparts" />
			</Helmet>

			<section id="SpareParts">
				<h1 className={utils.sectionHeading}>Spare-Parts</h1>
				{isAdmin && suppliers && <SparepartForm suppliers={suppliers} />}

				<div className={utils.optionsBtn}>
					<WarrantyForm />
					<RegistrationForm />
				</div>

				{suppliers && (
					<div className={utils.index}>
						<h1 className={utils.indexHeading}>Suppliers</h1>
						{suppliers.map(link => (
							<a key={link.name} href={`#${link.name}`}>
								<h1 className={utils.indexItem}>{link.name}</h1>
							</a>
						))}
					</div>
				)}
				{suppliers.map(supplier => (
					<div className={styles.supplierCard} key={supplier.id} id={supplier.name}>
						<h1 className={utils.mainHeading}>{supplier.name}</h1>
						<img
							src={supplier.logo_image ?? '/default.jpg'}
							className={styles.partsImage}
							alt="Supplier logo"
							width={200}
							height={200}
						/>
						<button className={utils.btn}>
							<Link to={`/spareparts/${supplier.id}`}>
								Spare-Parts
								<FontAwesomeIcon icon={faRightToBracket} />
							</Link>
						</button>
					</div>
				))}
			</section>
		</>
	);
};

export default SpareParts;
