package com.panty.test.springboottest1utils;


import com.panty.test.springboottest1utils.helper.TimeHelper;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.Date;

@RunWith(SpringRunner.class)
@SpringBootTest(classes = SpringbootTest1UtilsApplication.class)
public class MyJunitTest {
    @Test
    public void test() {
        TimeHelper timeHelper = new TimeHelper();
        Date date = timeHelper.getCurDate();
        System.out.println("test result ：" + date.toString());
    }
}
