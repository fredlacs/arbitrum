/* Generated by ts-generator ver. 0.0.8 */
/* tslint:disable */

import { Contract, ContractTransaction, EventFilter, Signer } from 'ethers'
import { Listener, Provider } from 'ethers/providers'
import { Arrayish, BigNumber, BigNumberish, Interface } from 'ethers/utils'
import {
  TransactionOverrides,
  TypedEventDescription,
  TypedFunctionDescription,
} from '.'

interface GlobalInboxInterface extends Interface {
  functions: {
    getERC20Balance: TypedFunctionDescription<{
      encode([_tokenContract, _owner]: [string, string]): string
    }>

    getERC721Tokens: TypedFunctionDescription<{
      encode([_erc721, _owner]: [string, string]): string
    }>

    getEthBalance: TypedFunctionDescription<{
      encode([_owner]: [string]): string
    }>

    getPaymentOwner: TypedFunctionDescription<{
      encode([originalOwner, nodeHash, messageIndex]: [
        string,
        Arrayish,
        BigNumberish
      ]): string
    }>

    hasERC721: TypedFunctionDescription<{
      encode([_erc721, _owner, _tokenId]: [
        string,
        string,
        BigNumberish
      ]): string
    }>

    ownedERC20s: TypedFunctionDescription<{
      encode([_owner]: [string]): string
    }>

    ownedERC721s: TypedFunctionDescription<{
      encode([_owner]: [string]): string
    }>

    transferPayment: TypedFunctionDescription<{
      encode([originalOwner, newOwner, nodeHash, messageIndex]: [
        string,
        string,
        Arrayish,
        BigNumberish
      ]): string
    }>

    withdrawERC20: TypedFunctionDescription<{
      encode([_tokenContract]: [string]): string
    }>

    withdrawERC721: TypedFunctionDescription<{
      encode([_erc721, _tokenId]: [string, BigNumberish]): string
    }>

    withdrawEth: TypedFunctionDescription<{ encode([]: []): string }>

    getInbox: TypedFunctionDescription<{ encode([account]: [string]): string }>

    sendMessages: TypedFunctionDescription<{
      encode([_messages, messageCounts, nodeHashes]: [
        Arrayish,
        BigNumberish[],
        Arrayish[]
      ]): string
    }>

    sendTransactionMessage: TypedFunctionDescription<{
      encode([_chain, _to, _seqNumber, _value, _data]: [
        string,
        string,
        BigNumberish,
        BigNumberish,
        Arrayish
      ]): string
    }>

    depositEthMessage: TypedFunctionDescription<{
      encode([_chain, _to]: [string, string]): string
    }>

    depositERC20Message: TypedFunctionDescription<{
      encode([_chain, _to, _erc20, _value]: [
        string,
        string,
        string,
        BigNumberish
      ]): string
    }>

    depositERC721Message: TypedFunctionDescription<{
      encode([_chain, _to, _erc721, _id]: [
        string,
        string,
        string,
        BigNumberish
      ]): string
    }>

    forwardContractTransactionMessage: TypedFunctionDescription<{
      encode([_to, _from, _value, _data]: [
        string,
        string,
        BigNumberish,
        Arrayish
      ]): string
    }>

    forwardEthMessage: TypedFunctionDescription<{
      encode([_to, _from]: [string, string]): string
    }>

    deliverTransactionBatch: TypedFunctionDescription<{
      encode([chain, transactions]: [string, Arrayish]): string
    }>
  }

  events: {
    ContractTransactionMessageDelivered: TypedEventDescription<{
      encodeTopics([chain, to, from, value, data, messageNum]: [
        string | null,
        string | null,
        string | null,
        null,
        null,
        null
      ]): string[]
    }>

    ERC20DepositMessageDelivered: TypedEventDescription<{
      encodeTopics([chain, to, from, erc20, value, messageNum]: [
        string | null,
        string | null,
        string | null,
        null,
        null,
        null
      ]): string[]
    }>

    ERC721DepositMessageDelivered: TypedEventDescription<{
      encodeTopics([chain, to, from, erc721, id, messageNum]: [
        string | null,
        string | null,
        string | null,
        null,
        null,
        null
      ]): string[]
    }>

    EthDepositMessageDelivered: TypedEventDescription<{
      encodeTopics([chain, to, from, value, messageNum]: [
        string | null,
        string | null,
        string | null,
        null,
        null
      ]): string[]
    }>

    PaymentTransfer: TypedEventDescription<{
      encodeTopics([
        nodeHash,
        messageIndex,
        originalOwner,
        prevOwner,
        newOwner,
      ]: [null, null, null, null, null]): string[]
    }>

    TransactionMessageBatchDelivered: TypedEventDescription<{
      encodeTopics([chain]: [string | null]): string[]
    }>

    TransactionMessageDelivered: TypedEventDescription<{
      encodeTopics([chain, to, from, seqNumber, value, data]: [
        string | null,
        string | null,
        string | null,
        null,
        null,
        null
      ]): string[]
    }>
  }
}

export class GlobalInbox extends Contract {
  connect(signerOrProvider: Signer | Provider | string): GlobalInbox
  attach(addressOrName: string): GlobalInbox
  deployed(): Promise<GlobalInbox>

  on(event: EventFilter | string, listener: Listener): GlobalInbox
  once(event: EventFilter | string, listener: Listener): GlobalInbox
  addListener(eventName: EventFilter | string, listener: Listener): GlobalInbox
  removeAllListeners(eventName: EventFilter | string): GlobalInbox
  removeListener(eventName: any, listener: Listener): GlobalInbox

  interface: GlobalInboxInterface

  functions: {
    getERC20Balance(_tokenContract: string, _owner: string): Promise<BigNumber>

    getERC721Tokens(_erc721: string, _owner: string): Promise<BigNumber[]>

    getEthBalance(_owner: string): Promise<BigNumber>

    getPaymentOwner(
      originalOwner: string,
      nodeHash: Arrayish,
      messageIndex: BigNumberish
    ): Promise<string>

    hasERC721(
      _erc721: string,
      _owner: string,
      _tokenId: BigNumberish
    ): Promise<boolean>

    ownedERC20s(_owner: string): Promise<string[]>

    ownedERC721s(_owner: string): Promise<string[]>

    transferPayment(
      originalOwner: string,
      newOwner: string,
      nodeHash: Arrayish,
      messageIndex: BigNumberish,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    withdrawERC20(
      _tokenContract: string,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    withdrawERC721(
      _erc721: string,
      _tokenId: BigNumberish,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    withdrawEth(overrides?: TransactionOverrides): Promise<ContractTransaction>

    getInbox(
      account: string
    ): Promise<{
      0: string
      1: BigNumber
    }>

    sendMessages(
      _messages: Arrayish,
      messageCounts: BigNumberish[],
      nodeHashes: Arrayish[],
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    sendTransactionMessage(
      _chain: string,
      _to: string,
      _seqNumber: BigNumberish,
      _value: BigNumberish,
      _data: Arrayish,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    depositEthMessage(
      _chain: string,
      _to: string,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    depositERC20Message(
      _chain: string,
      _to: string,
      _erc20: string,
      _value: BigNumberish,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    depositERC721Message(
      _chain: string,
      _to: string,
      _erc721: string,
      _id: BigNumberish,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    forwardContractTransactionMessage(
      _to: string,
      _from: string,
      _value: BigNumberish,
      _data: Arrayish,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    forwardEthMessage(
      _to: string,
      _from: string,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>

    deliverTransactionBatch(
      chain: string,
      transactions: Arrayish,
      overrides?: TransactionOverrides
    ): Promise<ContractTransaction>
  }

  getERC20Balance(_tokenContract: string, _owner: string): Promise<BigNumber>

  getERC721Tokens(_erc721: string, _owner: string): Promise<BigNumber[]>

  getEthBalance(_owner: string): Promise<BigNumber>

  getPaymentOwner(
    originalOwner: string,
    nodeHash: Arrayish,
    messageIndex: BigNumberish
  ): Promise<string>

  hasERC721(
    _erc721: string,
    _owner: string,
    _tokenId: BigNumberish
  ): Promise<boolean>

  ownedERC20s(_owner: string): Promise<string[]>

  ownedERC721s(_owner: string): Promise<string[]>

  transferPayment(
    originalOwner: string,
    newOwner: string,
    nodeHash: Arrayish,
    messageIndex: BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  withdrawERC20(
    _tokenContract: string,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  withdrawERC721(
    _erc721: string,
    _tokenId: BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  withdrawEth(overrides?: TransactionOverrides): Promise<ContractTransaction>

  getInbox(
    account: string
  ): Promise<{
    0: string
    1: BigNumber
  }>

  sendMessages(
    _messages: Arrayish,
    messageCounts: BigNumberish[],
    nodeHashes: Arrayish[],
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  sendTransactionMessage(
    _chain: string,
    _to: string,
    _seqNumber: BigNumberish,
    _value: BigNumberish,
    _data: Arrayish,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  depositEthMessage(
    _chain: string,
    _to: string,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  depositERC20Message(
    _chain: string,
    _to: string,
    _erc20: string,
    _value: BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  depositERC721Message(
    _chain: string,
    _to: string,
    _erc721: string,
    _id: BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  forwardContractTransactionMessage(
    _to: string,
    _from: string,
    _value: BigNumberish,
    _data: Arrayish,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  forwardEthMessage(
    _to: string,
    _from: string,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  deliverTransactionBatch(
    chain: string,
    transactions: Arrayish,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction>

  filters: {
    ContractTransactionMessageDelivered(
      chain: string | null,
      to: string | null,
      from: string | null,
      value: null,
      data: null,
      messageNum: null
    ): EventFilter

    ERC20DepositMessageDelivered(
      chain: string | null,
      to: string | null,
      from: string | null,
      erc20: null,
      value: null,
      messageNum: null
    ): EventFilter

    ERC721DepositMessageDelivered(
      chain: string | null,
      to: string | null,
      from: string | null,
      erc721: null,
      id: null,
      messageNum: null
    ): EventFilter

    EthDepositMessageDelivered(
      chain: string | null,
      to: string | null,
      from: string | null,
      value: null,
      messageNum: null
    ): EventFilter

    PaymentTransfer(
      nodeHash: null,
      messageIndex: null,
      originalOwner: null,
      prevOwner: null,
      newOwner: null
    ): EventFilter

    TransactionMessageBatchDelivered(chain: string | null): EventFilter

    TransactionMessageDelivered(
      chain: string | null,
      to: string | null,
      from: string | null,
      seqNumber: null,
      value: null,
      data: null
    ): EventFilter
  }

  estimate: {
    getERC20Balance(_tokenContract: string, _owner: string): Promise<BigNumber>

    getERC721Tokens(_erc721: string, _owner: string): Promise<BigNumber>

    getEthBalance(_owner: string): Promise<BigNumber>

    getPaymentOwner(
      originalOwner: string,
      nodeHash: Arrayish,
      messageIndex: BigNumberish
    ): Promise<BigNumber>

    hasERC721(
      _erc721: string,
      _owner: string,
      _tokenId: BigNumberish
    ): Promise<BigNumber>

    ownedERC20s(_owner: string): Promise<BigNumber>

    ownedERC721s(_owner: string): Promise<BigNumber>

    transferPayment(
      originalOwner: string,
      newOwner: string,
      nodeHash: Arrayish,
      messageIndex: BigNumberish
    ): Promise<BigNumber>

    withdrawERC20(_tokenContract: string): Promise<BigNumber>

    withdrawERC721(_erc721: string, _tokenId: BigNumberish): Promise<BigNumber>

    withdrawEth(): Promise<BigNumber>

    getInbox(account: string): Promise<BigNumber>

    sendMessages(
      _messages: Arrayish,
      messageCounts: BigNumberish[],
      nodeHashes: Arrayish[]
    ): Promise<BigNumber>

    sendTransactionMessage(
      _chain: string,
      _to: string,
      _seqNumber: BigNumberish,
      _value: BigNumberish,
      _data: Arrayish
    ): Promise<BigNumber>

    depositEthMessage(_chain: string, _to: string): Promise<BigNumber>

    depositERC20Message(
      _chain: string,
      _to: string,
      _erc20: string,
      _value: BigNumberish
    ): Promise<BigNumber>

    depositERC721Message(
      _chain: string,
      _to: string,
      _erc721: string,
      _id: BigNumberish
    ): Promise<BigNumber>

    forwardContractTransactionMessage(
      _to: string,
      _from: string,
      _value: BigNumberish,
      _data: Arrayish
    ): Promise<BigNumber>

    forwardEthMessage(_to: string, _from: string): Promise<BigNumber>

    deliverTransactionBatch(
      chain: string,
      transactions: Arrayish
    ): Promise<BigNumber>
  }
}
