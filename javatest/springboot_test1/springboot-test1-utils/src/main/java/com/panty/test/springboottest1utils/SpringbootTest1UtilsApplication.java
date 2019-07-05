package com.panty.test.springboottest1utils;

import com.panty.test.springboottest1utils.dubbotest.DubboInvokeUtil;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class SpringbootTest1UtilsApplication {

    public static void main(String[] args) {

        SpringApplication.run(SpringbootTest1UtilsApplication.class, args);

        DubboInvokeUtil dubboInvokeUtil =new DubboInvokeUtil();
        dubboInvokeUtil.invokeDubbo(new Object[]{"xxx"});
    }
}
