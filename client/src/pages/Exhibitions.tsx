import utils from '../styles/Utils.module.css';
import styles from '../styles/Blogs.module.css';

import { Exhibition } from '../types/blogTypes';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';

import ErrorPage from '../layouts/Error';
import Loading from '../layouts/Loading';
import ExhibitionForm from '../forms/ExhibitionForm';
import DeleteButton from '../components/DeleteButton';
import { Helmet } from 'react-helmet';
import {FC} from "react";

const Exhibitions: FC = () => {
	const { isAdmin } = useUserStore();
	const { data: exhibitions, isLoading, isError } = useGetResource<Exhibition[]>('exhibitions');

	if (isError) return <ErrorPage />;
	if (isLoading) return <Loading />;

	return (
		<>
			<Helmet>
				<title>Exhibitions - Farmec Ireland</title>
				<meta name="description" content="Check out the upcoming exhibitions where Farmec will showcase its latest machinery. Don't miss the chance to see our products in action." />

				<meta property="og:title" content="Exhibitions - Farmec Ireland" />
				<meta property="og:description" content="Check out the upcoming exhibitions where Farmec will showcase its latest machinery. Don't miss the chance to see our products in action." />
				<meta property="og:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />
				<meta property="og:url" content="https://www.farmec.ie/exhibitions" />
				<meta property="og:type" content="website" />

				<meta name="twitter:card" content="summary_large_image" />
				<meta name="twitter:title" content="Exhibitions - Farmec Ireland" />
				<meta name="twitter:description" content="Check out the upcoming exhibitions where Farmec will showcase its latest machinery. Don't miss the chance to see our products in action." />
				<meta name="twitter:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />

				<link rel="canonical" href="https://www.farmec.ie/exhibitions" />
			</Helmet>

			<section id="exhibitions">
				<h1 className={utils.sectionHeading}>Exhibitions:</h1>
				<h1 className={utils.subHeading}>Check out upcoming events related to Farmec</h1>
				{isAdmin && <ExhibitionForm />}

				{exhibitions && (
					<div className={styles.exhibitions}>
						{exhibitions.map(exhibition => (
							<div className={styles.exhibition} key={exhibition.id}>
								<h1 className={utils.mainHeading}>{exhibition.title}</h1>
								<p className={utils.paragraph}>{exhibition.date}</p>
								<p className={utils.paragraph}>{exhibition.location}</p>
								<p className={utils.paragraph}>{exhibition.info}</p>
								{isAdmin && exhibition.id && (
									<div className={utils.optionsBtn}>
										<ExhibitionForm id={exhibition.id} exhibition={exhibition} />
										<DeleteButton id={exhibition.id} resourceKey="exhibitions" />
									</div>
								)}
							</div>
						))}
					</div>
				)}
			</section>
		</>
	);
};

export default Exhibitions;
