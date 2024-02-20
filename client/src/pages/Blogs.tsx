import utils from '../styles/Utils.module.css';
import styles from '../styles/Blogs.module.css';

import ErrorPage from '../layouts/Error';
import Loading from '../layouts/Loading';
import BlogForm from '../forms/BlogForm';
import DeleteButton from '../components/DeleteButton';

import { Link } from 'react-router-dom';
import { Blog } from '../types/blogTypes';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faRightToBracket } from '@fortawesome/free-solid-svg-icons/faRightToBracket';

const Blogs: React.FC = () => {
    const { isAdmin } = useUserStore();
    const { data: blogs, isLoading, isError } = useGetResource<Blog[]>('blogs');

    if (isError) return <ErrorPage />;
    if (isLoading) return <Loading />;

    return (
        <section id="blog">
            <h1 className={utils.sectionHeading}>Check out our Latest Blog Posts</h1>
            <p className={utils.subHeading}> Read our latest news</p>
            {blogs && (
                <div className={utils.index}>
                    <h1 className={utils.indexHeading}>Suppliers</h1>
                    {blogs.map(link => (
                        <a key={link.title} href={`#${link.title}`}>
                            <h1 className={utils.indexItem}>{link.title}</h1>
                        </a>
                    ))}
                </div>
            )}

            {blogs?.map(blog => (
                <div className={styles.blogGrid} key={blog.id} id={blog.title || ''}>
                    <div className={styles.blogCard}>
                        <img
                            className={styles.blogImage}
                            src={blog.main_image || '/default.jpg'}
                            alt={'/default.jpg'}
                            width={300}
                            height={300}
                        />
                        <div className={styles.blogLink}>
                            <h1 className={utils.mainHeading}>{blog.title}</h1>
                            <p className={utils.paragraph}>{blog.subheading}</p>
                            <p className={utils.paragraph}>{blog.body}</p>
                            <button className={utils.btnForm}>
                                <Link to={`/blogs/${blog.id}`}>
                                    Read More
                                    <FontAwesomeIcon icon={faRightToBracket} />
                                </Link>
                            </button>
                        </div>
                    </div>
                    {isAdmin && blog.id && (
                        <div className={utils.optionsBtn}>
                            <BlogForm id={blog.id} blog={blog} />
                            <DeleteButton id={blog.id} resourceKey="blogs" />
                        </div>
                    )}
                </div>
            ))}
            {isAdmin && <BlogForm />}
        </section>
    );
};

export default Blogs;
