package com.emarket.orderservice.services;

import com.emarket.orderservice.models.Order;
import org.springframework.stereotype.Component;

import java.util.UUID;

@Component
public class OrderService {


    public Order createOrder(Order order){

        // persist
        Order order1 = new Order();
        order1.setId(UUID.randomUUID().toString());
        return order1;
    }

    public Order getOrder(String id){
        // get from db
        Order order = new Order();
        order.setId(id);
        return order;
    }
}
