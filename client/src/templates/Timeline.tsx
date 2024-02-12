import styles from '../styles/About.module.css';

import { Timeline } from '../types/aboutTypes';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';

import Error from '../layouts/Error';
import Loading from '../layouts/Loading';
import TimelineCard from '../components/TimelineCard';

const Timelines: React.FC = () => {
    const { isAdmin } = useUserStore();
    const { data: timelines, isLoading, isError } = useGetResource<Timeline[]>('timelines');

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

    return (
        <section id="timeline">
            <div className={styles.timeline}>
                {timelines?.map((timeline: Timeline) => (
                    <TimelineCard timeline={timeline} isAdmin={isAdmin} key={timeline.id} />
                ))}
            </div>
        </section>
    );
};

export default Timelines;
