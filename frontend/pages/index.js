import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'

export default function Home() {

 const generateInvestigator = async event => {
    event.preventDefault()

    const res = await fetch('http://localhost:8080/v1/investigator', {
      body: JSON.stringify({
        name: event.target.name.value,
        age: event.target.age.value,
        residence: event.target.residence.value,
        birthplace: event.target.birthplace.value,
        occupation: event.target.occupation.value,
      }),

      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST'
    })

    const result = await res.json()
	 alert(`Is this your full name: ${result.investigator.name}`)
  }


  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Welcome to <a href="https://nextjs.org">Next.js!</a>
        </h1>

        <p className={styles.description}>
          Get started by editing{' '}
          <code className={styles.code}>pages/index.js</code>
        </p>

	 <form onSubmit={generateInvestigator}>
     	<label htmlFor="name">Name</label>
      		<input id="name" name="name" type="text" required />
	 
		<label htmlFor="age">Age</label>
      		<input id="age" name="age" type="number" required />
       
		<label htmlFor="residence">Residence</label>
      		<input id="residence" name="residence" type="text" required />
       
		 <label htmlFor="birthplace">Birthplace</label>
      		<input id="birthplace" name="birthplace" type="text"  required />

		 <label htmlFor="occupation">Occupation</label>
      		<input id="occupation" name="occupation" type="text" required />
                
		 <button type="submit">Generate</button>
    </form>
     </main>

      <footer className={styles.footer}>
        <a
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by{' '}
          <span className={styles.logo}>
            <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16} />
          </span>
        </a>
      </footer>
    </div>
  )
}