import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Spareparts.module.css';

import { Link } from 'react-router-dom';
import { useParams } from 'react-router-dom';
import { Sparepart } from '../../types/supplierTypes';
import { useGetResourceById } from '../../hooks/genericHooks';
import { useSupplierStore, useUserStore } from '../../lib/store';

import ErrorPage from '../../layouts/Error';
import Loading from '../../layouts/Loading';
import SparepartForm from '../../forms/SparePartsForm';
import DeleteButton from '../../components/DeleteButton';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faRightToBracket } from '@fortawesome/free-solid-svg-icons/faRightToBracket';
import { useEffect } from 'react';

const PartsDetail: React.FC = () => {
	const { isAdmin } = useUserStore();
	const { suppliers } = useSupplierStore();

	const id = useParams<{ id: string }>().id as string;
	const { data: spareparts, isLoading, isError, error } = useGetResourceById<Sparepart[]>('spareparts', id);

	useEffect(() => {}, [id]);

	if (isError) return <ErrorPage />;
	if (isLoading) return <Loading />;

	const imageError = (event: React.SyntheticEvent<HTMLImageElement, Event>) => {
		event.currentTarget.src = '/default.jpg';
	};

	if (!spareparts) {
		return (
			<div>
				<h1 className={utils.sectionHeading}>No spare parts available for this supplier</h1>
				<button className={utils.btn}>
					<Link to={'/spareparts'}>
						Spare Parts <FontAwesomeIcon icon={faRightToBracket} />
					</Link>
				</button>
			</div>
		);
	}

	return (
		<section id="partsDetail">
			<h1 className={utils.sectionHeading}>Parts Catalogues</h1>
			{isAdmin && suppliers && <SparepartForm suppliers={suppliers} />}

			{spareparts && (
				<div className={utils.index}>
					<h1 className={utils.indexHeading}>spareparts</h1>
					{spareparts.map(link => (
						<a key={link.name} href={`#${link.name}`}>
							<h1 className={utils.indexItem}>{link.name}</h1>
						</a>
					))}
				</div>
			)}
			<div className={styles.partsColumn}>
				{spareparts ? (
					spareparts.map(sparepart => (
						<div className={styles.sparepartGrid} key={sparepart.id}>
							<div className={styles.sparepartsCard} id={sparepart.name || ''}>
								<div className={styles.sparepartsGrid}>
									<div className={styles.sparepartsInfo}>
										<h1 className={utils.mainHeading}>{sparepart.name}</h1>
										<button className={utils.btn}>
											<Link to={sparepart.spare_parts_link || '#'} target="_blank">
												Parts Catalogue <FontAwesomeIcon icon={faRightToBracket} />
											</Link>
										</button>
									</div>
									<img
										src={sparepart.parts_image}
										alt={'/default.jpg'}
										className={styles.sparepartsLogo}
										width={600}
										height={600}
										onError={imageError}
									/>
								</div>
							</div>
							{isAdmin && suppliers && sparepart.id && (
								<div className={utils.optionsBtn}>
									<SparepartForm suppliers={suppliers} sparepart={sparepart} id={sparepart.id} />
									<DeleteButton id={sparepart.id} resourceKey={'spareparts'} />
								</div>
							)}
						</div>
					))
				) : (
					<div>error: {error || 'Unknown error'}</div>
				)}
			</div>
		</section>
	);
};

export default PartsDetail;
