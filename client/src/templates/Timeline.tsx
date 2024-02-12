import styles from '../styles/About.module.css';

import { useGetResource } from '../hooks/genericHooks';
import { useUserStore } from '../lib/store';
import { Timeline } from '../types/aboutTypes';

import TimelineCard from '../components/TimelineCard';

const Timelines: React.FC = () => {
    const { isAdmin } = useUserStore();

    const { data: timelines, isLoading, error } = useGetResource<Timeline[]>('timelines');

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

    return (
        <section id="timeline">
            <div className={styles.timeline}>
                {timelines ? (
                    timelines.map((timeline: Timeline) => (
                        <TimelineCard timeline={timeline} isAdmin={isAdmin} key={timeline.id} />
                    ))
                ) : (
                    <div>error: {error?.message || 'Unknown error'}</div>
                )}
            </div>
        </section>
    );
};

export default Timelines;
