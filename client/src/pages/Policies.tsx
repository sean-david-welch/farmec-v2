import utils from '../styles/Utils.module.css';
import styles from '../styles/About.module.css';

import {Resources} from '../types/dataTypes';
import {useUserStore} from '../lib/store';
import {Privacy, Terms} from '../types/aboutTypes';
import {useMultipleResourcesWithoutId} from '../hooks/genericHooks';

import ErrorPage from '../layouts/Error';
import Loading from '../layouts/Loading';
import TermForm from '../forms/TermForm';
import PrivacyForm from '../forms/PrivacyForm';
import DeleteButton from '../components/DeleteButton';
import {Helmet} from "react-helmet";
import {FC} from "react";

const Policies: FC = () => {
	const { isAdmin } = useUserStore();

	const resourceKeys: (keyof Resources)[] = ['terms', 'privacys'];
	const { data, isLoading, isError } = useMultipleResourcesWithoutId(resourceKeys);

	if (isError) return <ErrorPage />;
	if (isLoading) return <Loading />;

	const [terms, privacys] = data;

	return (
		<>
			<Helmet>
				<title>Policies - Farmec Ireland</title>
				<meta name="description" content="Read more about our Privacy Policy and how we use your data"/>

				<meta property="og:title" content="Policies - Farmec Ireland"/>
				<meta property="og:description" content="Discover Farmec's staff, history, and vision for the future."/>
				<meta property="og:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"/>
				<meta property="og:url" content="https://www.farmec.ie/policies"/>
				<meta property="og:type" content="website"/>

				<meta name="twitter:card" content="summary_large_image"/>
				<meta name="twitter:title" content="Policies - Farmec Ireland"/>
				<meta name="twitter:description" content="Read more about our Privacy Policy and how we use your data"/>
				<meta name="twitter:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"/>
				<link rel="canonical" href="https://www.farmec.ie/policies"/>
			</Helmet>
			<section id="policies">
				<div className={styles.terms}>
					<h1 className={utils.sectionHeading}>Terms of Service</h1>
					{isAdmin && <TermForm/>}

					{terms.map((term: Terms) => (
						<div className={styles.infoCard} key={term.id}>
							<h2 className={utils.mainHeading}>{term.title}</h2>
							<p className={utils.paragraph}>{term.body}</p>
							{isAdmin && term.id && (
								<div className={utils.optionsBtn}>
									<TermForm id={term.id} term={term}/>
									<DeleteButton id={term.id} resourceKey="terms"/>
								</div>
							)}
						</div>
					))}
				</div>

				<div className={styles.privacy}>
					<h1 className={utils.sectionHeading}>Privacy Policy</h1>
					{isAdmin && <PrivacyForm/>}

					{privacys.map((policy: Privacy) => (
						<div className={styles.infoCard} key={policy.id}>
							<h2 className={utils.mainHeading}>{policy.title}</h2>
							<p className={utils.paragraph}>{policy.body}</p>
							{isAdmin && policy.id && (
								<div className={utils.optionsBtn}>
									<PrivacyForm id={policy.id} privacy={policy}/>
									<DeleteButton id={policy.id} resourceKey="privacys"/>
								</div>
							)}
						</div>
					))}
				</div>
			</section>
		</>
	);
};

export default Policies;
