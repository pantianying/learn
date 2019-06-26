package com.panty.test.springboottest1api;

import com.panty.test.springboottest1utils.dubbotest.DubboInvokeUtil;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ImportResource;

@SpringBootApplication
@ImportResource(locations = {"classpath:spring-all.xml"})
public class SpringbootTest1ApiApplication {

    public static void main(String[] args) {
        SpringApplication.run(SpringbootTest1ApiApplication.class, args);

        DubboInvokeUtil dubboInvokeUtil = new DubboInvokeUtil();
        System.out.println("start invoke");
        dubboInvokeUtil.invokeDubbo(new String[]{"aa"});

    }

}
