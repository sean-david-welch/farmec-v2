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
import { Helmet } from 'react-helmet';

const Blogs: React.FC = () => {
	const { isAdmin } = useUserStore();
	const { data: blogs, isLoading, isError } = useGetResource<Blog[]>('blogs');

	if (isError) return <ErrorPage />;
	if (isLoading) return <Loading />;

	const imageError = (event: React.SyntheticEvent<HTMLImageElement, Event>) => {
		event.currentTarget.src = '/default.jpg';
	};

	return (
		<>
			<Helmet>
				<title>Latest Blog Posts - Farmec Ireland</title>
				<meta name="description" content="Check out the latest blog posts from Farmec Ireland. Stay up to date with our latest news and insights." />

				<meta property="og:title" content="Latest Blog Posts - Farmec Ireland" />
				<meta property="og:description" content="Check out the latest blog posts from Farmec Ireland. Stay up to date with our latest news and insights." />
				<meta property="og:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />
				<meta property="og:url" content="https://www.farmec.ie/blogs" />
				<meta property="og:type" content="website" />

				<meta name="twitter:card" content="summary_large_image" />
				<meta name="twitter:title" content="Latest Blog Posts - Farmec Ireland" />
				<meta name="twitter:description" content="Check out the latest blog posts from Farmec Ireland. Stay up to date with our latest news and insights." />
				<meta name="twitter:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />

				<link rel="canonical" href="https://www.farmec.ie/blogs" />
			</Helmet>

			<section id="blog">
				<h1 className={utils.sectionHeading}>Check out our Latest Blog Posts</h1>
				<p className={utils.subHeading}>Read our latest news</p>
				{isAdmin && <BlogForm />}
				{blogs && (
					<div className={utils.index}>
						<h1 className={utils.indexHeading}>Blogs</h1>
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
								src={blog.main_image}
								alt="Blog image"
								width={300}
								height={300}
								onError={imageError}
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
			</section>
		</>
	);
};

export default Blogs;
