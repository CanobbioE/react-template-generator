package main

const componentTemplate = `
import React from "react";
import { View, Text} from "react-native";

interface {{.Name}}Props {

}

const {{.Name}} = (props: {{.Name}}Props) => {
    return(
        <View>
            <Text>Hello World!</Text>
        </View>
    )
}

export default {{.Name}}; 
`

const styleTemplate = `
import { StyleSheet } from "react-native";
import { fonts } from "../../Styles";

const styles = StyleSheet.create({
  
});

export default styles; 
`
