import Head from 'next/head'
import Image from 'next/image'
import Login from '../components/login'
import Signup from '../components/signup'
import puplepattern from 'src/asset/texture-dark-background-purple-3840x2715-3086.jpg'
import React from 'react'

class Home extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      isLogin: true
    }

    this.handler = this.handler.bind(this)
  }

  handler() {
    this.setState(prevState => ({
      isLogin: !prevState.isLogin
    }));
  }


  render() {
    return (
      <>
      <Head>
        <title>Sirkelin</title>
        <meta name="description" content="Sirkelin: " />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className=''>
        <div className='grid h-screen place-items-center' >
            <div className='bg-slate-900 rounded-xl shadow-2xl flex wh-l '>
                <Image src={puplepattern} className='wh-i rounded-xl' />
                { this.state.isLogin ? <Login handler={this.handler} /> : <Signup handler={this.handler} /> }
            </div>
        </div>
      </main>
    </>
  )}
}


export default Home;