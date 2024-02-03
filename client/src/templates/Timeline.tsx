import styles from '../styles/About.module.css';

import { useGetResource } from '../hooks/genericHooks';
import { useUserStore } from '../lib/store';
import { Timeline } from '../types/aboutTypes';

import TimelineCard from '../components/TimlineCard';

const Timelines: React.FC = () => {
    const { isAdmin } = useUserStore();

    const { data: timelines, isLoading, error } = useGetResource<Timeline[]>('timelines');

    if (isLoading) {
        return <div>Loading...</div>;
    }

    return (
        <section id="timeline">
            <div className={styles.timeline}>
                {timelines ? (
                    timelines.map((timeline: Timeline) => (
                        <TimelineCard timeline={timeline} isAdmin={isAdmin} />
                    ))
                ) : (
                    <div>error: {error?.message || 'Unknown error'}</div>
                )}
            </div>
        </section>
    );
};

export default Timelines;
