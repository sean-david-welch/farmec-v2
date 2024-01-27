import styles from '../styles/Suppliers.module.css';

import { Link } from 'react-router-dom';

interface SocialLinksProps {
    facebook?: string | null;
    twitter?: string | null;
    instagram?: string | null;
    linkedin?: string | null;
    website?: string | null;
    youtube?: string | null;
}

export const SocialLinks = ({ facebook, twitter, instagram, linkedin, website, youtube }: SocialLinksProps) => {
    const facebookIcon = '/icons/facebook.svg';
    const twitterIcon = '/icons/twitter.svg';
    const instagramIcon = '/icons/instagram.svg';
    const linkedinIcon = '/icons/linkedin.svg';
    const websiteIcon = '/icons/website.svg';
    const youtubeIcon = '/icons/youtube.svg';

    return (
        <div className={styles.socialLinks}>
            {facebook && (
                <Link to={facebook} target="_blank" className={styles.facebookButton}>
                    <img src={facebookIcon} className={styles.icon} alt="Facebook" />
                </Link>
            )}
            {twitter && (
                <Link to={twitter} target="_blank" className={styles.twitterButton}>
                    <img src={twitterIcon} className={styles.icon} alt="Twitter" />
                </Link>
            )}
            {instagram && (
                <Link to={instagram} target="_blank" className={styles.instagramButton}>
                    <img src={instagramIcon} className={styles.icon} alt="Instagram" />
                </Link>
            )}
            {linkedin && (
                <Link to={linkedin} target="_blank" className={styles.linkedinButton}>
                    <img src={linkedinIcon} className={styles.icon} alt="LinkedIn" />
                </Link>
            )}
            {website && (
                <Link to={website} target="_blank" className={styles.websiteButton}>
                    <img src={websiteIcon} className={styles.icon} alt="Website" />
                </Link>
            )}
            {youtube && (
                <Link to={youtube} target="_blank" className={styles.youtubeButton}>
                    <img src={youtubeIcon} className={styles.icon} alt="YouTube" />
                </Link>
            )}
        </div>
    );
};
