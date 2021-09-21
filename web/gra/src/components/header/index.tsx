import React, {ReactElement, FC} from "react";
import { Button } from 'antd';

const Header: FC = (): ReactElement => {

    return(
        <div>
            <Button type="primary">Button</Button>
        </div>
    )
}

export default Header;