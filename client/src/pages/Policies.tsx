import utils from '../styles/Utils.module.css';
import styles from '../styles/About.module.css';

import { Resources } from '../types/dataTypes';
import { useUserStore } from '../lib/store';
import { Privacy, Terms } from '../types/aboutTypes';
import { useMultipleResourcesWithoutId } from '../hooks/genericHooks';

import Error from '../layouts/Error';
import Loading from '../layouts/Loading';
import TermForm from '../forms/TermForm';
import PrivacyForm from '../forms/PrivacyForm';
import DeleteButton from '../components/DeleteButton';

const Policies: React.FC = () => {
    const { isAdmin } = useUserStore();

    const resourceKeys: (keyof Resources)[] = ['terms', 'privacys'];
    const { data, isLoading, isError } = useMultipleResourcesWithoutId(resourceKeys);

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

    const [terms, privacys] = data;

    return (
        <section id="policies">
            <div className={styles.terms}>
                <h1 className={utils.sectionHeading}>Terms of Service</h1>
                {terms.map((term: Terms) => (
                    <div className={styles.infoCard} key={term.id}>
                        <h2 className={utils.mainHeading}>{term.title}</h2>
                        <p className={utils.paragraph}>{term.body}</p>
                        {isAdmin && term.id && (
                            <div className={utils.optionsBtn}>
                                <TermForm id={term.id} term={term} />
                                <DeleteButton id={term.id} resourceKey="terms" />
                            </div>
                        )}
                    </div>
                ))}

                {isAdmin && <TermForm />}
            </div>

            <div className={styles.privacy}>
                <h1 className={utils.sectionHeading}>Privacy Policy</h1>
                {privacys.map((policy: Privacy) => (
                    <div className={styles.infoCard} key={policy.id}>
                        <h2 className={utils.mainHeading}>{policy.title}</h2>
                        <p className={utils.paragraph}>{policy.body}</p>
                        {isAdmin && policy.id && (
                            <div className={utils.optionsBtn}>
                                <PrivacyForm id={policy.id} privacy={policy} />
                                <DeleteButton id={policy.id} resourceKey="privacys" />
                            </div>
                        )}
                    </div>
                ))}

                {isAdmin && <PrivacyForm />}
            </div>
        </section>
    );
};

export default Policies;
