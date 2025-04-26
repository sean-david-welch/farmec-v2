import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Suppliers.module.css';

import ErrorPage from '../../layouts/Error';
import Videos from '../../templates/Videos';
import Loading from '../../layouts/Loading';
import Machines from '../../templates/Machines';

import { useParams } from 'react-router-dom';
import { Resources } from '../../types/dataTypes';
import { useMultipleResourcesSlug} from '../../hooks/genericHooks';
import { useSupplierStore, useUserStore } from '../../lib/store';
import {useEffect, Fragment, FC} from 'react';
import SupplierForm from '../../forms/SupplierForm';
import DeleteButton from '../../components/DeleteButton';
import {Helmet} from "react-helmet";

const SuppliersDetails: FC = () => {
	const { isAdmin } = useUserStore();
	const { suppliers } = useSupplierStore();

	const slug = useParams<{ slug: string }>().slug as string;

	const resourceKeys: (keyof Resources)[] = ['suppliers', 'supplierMachine', 'videos'];
	const { data, isLoading, isError } = useMultipleResourcesSlug(slug, resourceKeys);

	useEffect(() => {}, [slug]);

	if (isError) return <ErrorPage />;
	if (isLoading) return <Loading />;

	const [supplier, machines, videos] = data;

	return (
		<>
			<Helmet>
				<title>{supplier ? `${supplier.name} - Farmec Ireland` : 'Supplier - Farmec Ireland'}</title>
				<meta name="description" content={supplier ? supplier.description : "Browse our Suppliers and learn more about the machines we offer."} />

				<meta property="og:title" content={supplier ? `${supplier.name} - Farmec Ireland` : 'Supplier - Farmec Ireland'} />
				<meta property="og:description" content={supplier ? supplier.description : "Browse our Suppliers and learn more about the machines we offer."} />
				<meta property="og:image" content={supplier?.marketing_image ? supplier.marketing_image : "https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"} />
				<meta property="og:url" content={`https://www.farmec.ie/suppliers/${supplier?.slug}`} />
				<meta property="og:type" content="website" />

				<meta name="twitter:card" content="summary_large_image" />
				<meta name="twitter:title" content={supplier ? `${supplier.name} - Farmec Ireland` : 'Supplier - Farmec Ireland'} />
				<meta name="twitter:description" content={supplier ? supplier.description : "Browse our Suppliers and learn more about the machines we offer."} />
				<meta name="twitter:image" content={supplier?.marketing_image ? supplier.marketing_image : "https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"} />

				<link rel="canonical" href={`https://www.farmec.ie/suppliers/${supplier?.slug}`} />
			</Helmet>
			<section id="supplierDetail">
				{supplier ? (
					<Fragment>
						<div className={styles.supplierHeading}>
							<h1 className={utils.sectionHeading}>{supplier.name}</h1>
							{isAdmin && supplier.id && (
								<div className={utils.optionsBtn}>
									<SupplierForm id={supplier.id} supplier={supplier}/>
									<DeleteButton id={supplier.id} resourceKey="suppliers"/>
								</div>
							)}
						</div>

						{machines && (
							<div className={utils.index}>
								<h1 className={utils.indexHeading}>Machines</h1>
								{machines.map((link: { name: string }) => (
									<a key={link.name} href={`#${link.name}`}>
										<h1 className={utils.indexItem}>{link.name}</h1>
									</a>
								))}
							</div>
						)}

						<div className={styles.supplierDetail}>
							<img
								src={supplier.marketing_image ?? '/default.jpg'}
								alt={'/dafault.jpg'}
								className={styles.supplierImage}
								width={750}
								height={750}
							/>

							<p className={styles.supplierDescription}>{supplier.description}</p>
						</div>
					</Fragment>
				) : null}

				{machines ? <Machines machines={machines} isAdmin={isAdmin}/> : null}
				{videos ? <Videos suppliers={suppliers} videos={videos} isAdmin={isAdmin}/> : null}
			</section>
		</>
	);
};

export default SuppliersDetails;
