import utils from '../../styles/Utils.module.css';

import ErrorPage from '../../layouts/Error';
import Loading from '../../layouts/Loading';
import Products from '../../templates/Products';

import {useParams} from 'react-router-dom';
import {Resources} from '../../types/dataTypes';
import {useUserStore} from '../../lib/store';
import {useMultipleResources} from '../../hooks/genericHooks';
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome';
import {faRightToBracket} from '@fortawesome/free-solid-svg-icons/faRightToBracket';
import {FC, useEffect} from 'react';
import ProductForm from '../../forms/ProductForm';
import {Helmet} from "react-helmet";

const MachineDetail: FC = () => {
	const { isAdmin } = useUserStore();
	const id = useParams<{ id: string }>().id as string;

	const resourceKeys: (keyof Resources)[] = ['machines', 'products'];
	const { data, isLoading, isError } = useMultipleResources(id, resourceKeys);

	useEffect(() => {}, [id]);

	if (isError) return <ErrorPage />;
	if (isLoading) return <Loading />;

	const [machine, products] = data;

	return (
		<>
			<Helmet>
				<title>{machine ? `${machine.name} - Farmec Ireland` : 'Machine - Farmec Ireland'}</title>
				<meta name="description" content={machine ? machine.description : "Browse our machines and products to learn more about what we offer."} />

				<meta property="og:title" content={machine ? `${machine.name} - Farmec Ireland` : 'Machine - Farmec Ireland'} />
				<meta property="og:description" content={machine ? machine.description : "Browse our machines and products to learn more about what we offer."} />
				<meta property="og:image" content={machine?.marketing_image ? machine.marketing_image : "https://www.farmec.ie/farmec_images/Machines/default-machine-image.webp"} />
				<meta property="og:url" content={`https://www.farmec.ie/machines/${machine?.id}`} />
				<meta property="og:type" content="website" />

				<meta name="twitter:card" content="summary_large_image" />
				<meta name="twitter:title" content={machine ? `${machine.name} - Farmec Ireland` : 'Machine - Farmec Ireland'} />
				<meta name="twitter:description" content={machine ? machine.description : "Browse our machines and products to learn more about what we offer."} />
				<meta name="twitter:image" content={machine?.marketing_image ? machine.marketing_image : "https://www.farmec.ie/farmec_images/Machines/default-machine-image.webp"} />

				<link rel="canonical" href={`https://www.farmec.ie/machines/${machine?.id}`} />
			</Helmet>

			<section id="machineDetail">
				<h1 className={utils.sectionHeading}>Products</h1>
				{isAdmin && <ProductForm machine={machine} />}
				{products && (
					<div className={utils.index}>
						<h1 className={utils.indexHeading}>Products:</h1>
						{products.map((link: { name: string }) => (
							<a key={link.name} href={`#${link.name}`}>
								<h1 className={utils.indexItem}>{link.name}</h1>
							</a>
						))}
						<button className={utils.btn}>
							<a href={machine.machine_link || '#'} target="_blank" rel="noopener noreferrer">
								Supplier Website
								<FontAwesomeIcon icon={faRightToBracket} />
							</a>
						</button>
					</div>
				)}
				{products && <Products id={id} products={products} isAdmin={isAdmin} />}
			</section>
		</>
	);
};

export default MachineDetail;
