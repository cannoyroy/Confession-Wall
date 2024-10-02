<template>
    <a :href="href" class="btn" :class="`btn-${color}`" @click.prevent="handleClick">
      <i :class="iconClass"></i>
    </a>
  </template>
  
  <script>
  export default {
    props: {
      href: {
        type: String,
        default: '#'
      },
      icon: {
        type: String,
        required: true
      },
      color: {
        type: String,
        default: 'default'
      }
    },
    computed: {
      iconClass() {
        return `fa ${this.icon}`;
      }
    },
    methods: {
      handleClick(event) {
        if (this.href && this.href.startsWith('#')) {
          this.$emit('click', event);
        } else if (this.href) {
          window.open(this.href, '_blank');
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .btn {
    display: inline-block;
    width: 60px;
    height: 60px;
    background: #f1f1f1;
    margin: 10px;
    border-radius: 30%;
    box-shadow: -5px 5px 15px -5px #f1f1f1;
    overflow: hidden;
    position: relative;
    transition: 0.3s linear;
    text-decoration: none;
  }
  
  .btn i {
    line-height: 60px;
    font-size: 26px;
    transition: 0.3s linear;
    color: #000; /* Default color */
  }
  
  .btn:before {
    content: "";
    position: absolute;
    width: 120%;
    height: 120%;
    transform: rotate(45deg);
    left: -110%;
    top: 90%;
  }
  
  .btn:hover {
    transform: scale(1.1);
  }
  
  .btn:hover i {
    transform: scale(1.2);
    color: #fff;
  }
  
  /* Social media colors */
  .btn-facebook:before {
    background: #3b5998;
  }
  
  .btn-twitter:before {
    background: #1da1f2;
  }
  
  .btn-instagram:before {
    background: #c32aa3;
  }
  
  .btn-google:before {
    background: #db4437;
  }
  </style>