// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(globalState) {
      return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(317)), function (_0_i0) {
        return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(906),false)).length)).multipliedBy(new BigNumber(619));
      })).length));
    };
    static fm3(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uypury"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(964)), function (_1_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), function (_2_i1) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("ceyuend")));
    };
    static fm9(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("atuia"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vwhjeu"), _dafny.Seq.UnicodeFromString("judictkq")));
    };
    static fm10(p0, p1, globalState) {
      return (_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(!(!(true))));
    };
    static fm11(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(611), new BigNumber(163))) {
          let _3_v0 = _compr_0;
          if (((new BigNumber(611)).isLessThanOrEqualTo(_3_v0)) && ((_3_v0).isLessThan(new BigNumber(163)))) {
            _coll0.push([_module.__default.safeDivisionInt(_3_v0, new BigNumber(-209)),true]);
          }
        }
        return _coll0;
      }()).length), new BigNumber(251))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(185)));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D2.create_DC3(_dafny.Seq.of(_module.D1.create_DC2(), _module.D1.create_DC2()));
      if (_source0.is_DC4) {
        if (false) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }
      } else if (_source0.is_DC3) {
        let _4___mcc_h0 = (_source0).cf2;
        let _5_cf2 = _4___mcc_h0;
        return new _dafny.CodePoint('l'.codePointAt(0));
      } else {
        let _6___mcc_h1 = (_source0).cf3;
        let _7_cf3 = _6___mcc_h1;
        if (true) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('b'.codePointAt(0));
        }
      }
    };
    static fm13(globalState) {
      return (function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lompq"))).Elements) {
          let _8_v0 = _compr_1;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lompq"))).contains(_8_v0)) {
            _coll1.add(_8_v0);
          }
        }
        return _coll1;
      }()).Intersect(((!(!(false))) ? (function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("i"), _dafny.Seq.UnicodeFromString("tyadwnh"), _dafny.Seq.UnicodeFromString("j"))).Elements) {
          let _9_v1 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("i"), _dafny.Seq.UnicodeFromString("tyadwnh"), _dafny.Seq.UnicodeFromString("j")), _9_v1)) {
            _coll2.add(_9_v1);
          }
        }
        return _coll2;
      }()) : (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("vbdruvjq")))));
    };
    static fm14(p0, p1, globalState) {
      return _dafny.Seq.of(!((_dafny.Set.fromElements(_module.D5.create_DC11(false, new BigNumber(456), false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-776)), function (_10_i0) {
  return new _dafny.CodePoint('j'.codePointAt(0));
}), new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))).length)))).IsDisjointFrom(_dafny.Set.fromElements(_module.D5.create_DC11(!(true), (_dafny.ZERO).minus(new BigNumber(-276)), false, _dafny.Seq.UnicodeFromString("sojj"), new BigNumber(456))))), !(false));
    };
    static fm15(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(253), new BigNumber(345), new BigNumber(535), new BigNumber((_dafny.Seq.UnicodeFromString("rpw")).length), new BigNumber(151)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-90)), function (_11_i0) {
        return new BigNumber(410);
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(660)), function (_12_i1) {
        return new BigNumber((_dafny.Seq.of(true, false)).length);
      }));
    };
    static fm16(p0, p1, p2, globalState) {
      return (new BigNumber(763)).isLessThan((_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber(-86))).multipliedBy(new BigNumber(292))));
    };
    static fm17(p0, p1, p2, globalState) {
      return _module.D1.create_DC2();
    };
    static fm18(p0, p1, p2, globalState) {
      return _module.D5.create_DC11((_dafny.Set.fromElements(false)).IsSubsetOf(_dafny.Set.fromElements(true)), new BigNumber((_dafny.Seq.of(false, true)).length), true, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ygxbvlc"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(645)), function (_13_i0) {
  return new _dafny.CodePoint('q'.codePointAt(0));
})), new BigNumber((function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(90)))).Elements) {
    let _14_v0 = _compr_3;
    if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(90)))).contains(_14_v0)) {
      _coll3.push([_module.__default.safeModuloInt(_14_v0, new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())),_dafny.Seq.UnicodeFromString("pvaooubs")]);
    }
  }
  return _coll3;
}()).length));
    };
    static fm19(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(752)), function (_15_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(536),new BigNumber(593));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(875)), function (_16_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(218),new BigNumber(506));
      }));
    };
    static fm20(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((false) === (false),(true) && (false));
    };
    static fm21(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-982)), function (_17_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(!(true),false);
      })) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-649)), function (_18_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(true,true);
      }))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,true)));
    };
    static fm22(globalState) {
      return _dafny.Set.fromElements(_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),true)).length),true)).length), new BigNumber((_dafny.Seq.of(new BigNumber(794), new BigNumber(181), new BigNumber(-949), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(103),false)).length))).length)), new BigNumber((((!(!(true))) ? (_dafny.Set.fromElements(false)) : (_dafny.Set.fromElements(true)))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length));
    };
    static fm23(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),(true) && (!(!(true))));
    };
    static fm24(p0, globalState) {
      return (function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(558)),false)).Keys.Elements) {
          let _19_v0 = _compr_4;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(558)),false)).contains(_19_v0)) {
            _coll4.add(_19_v0);
          }
        }
        return _coll4;
      }()).Union(_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber(245), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-52))).length), new BigNumber(818))).cardinality())), new BigNumber(584), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber(-743)), _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).length))));
    };
    static fm25(p0, globalState) {
      let _source1 = _module.D7.create_DC15(false);
      if (_source1.is_DC15) {
        let _20___mcc_h0 = (_source1).cf20;
        let _21_cf20 = _20___mcc_h0;
        return function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of ((_module.D13.create_DC29(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(15), new BigNumber(-640)), _dafny.Set.fromElements(new BigNumber(960))))).dtor_cf45).Elements) {
            let _22_v0 = _compr_5;
            if (_dafny.Seq.contains((_module.D13.create_DC29(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(15), new BigNumber(-640)), _dafny.Set.fromElements(new BigNumber(960))))).dtor_cf45, _22_v0)) {
              _coll5.push([_22_v0,_21_cf20]);
            }
          }
          return _coll5;
        }();
      } else if (_source1.is_DC14) {
        let _23___mcc_h1 = (_source1).cf19;
        let _24_cf19 = _23___mcc_h1;
        return function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), function (_25_i0) {
            return _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length));
          })).Elements) {
            let _26_v1 = _compr_6;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), function (_25_i0) {
              return _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length));
            }), _26_v1)) {
              _coll6.push([_26_v1,false]);
            }
          }
          return _coll6;
        }();
      } else {
        let _27___mcc_h2 = (_source1).cf21;
        let _28_cf21 = _27___mcc_h2;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(977)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(290), new BigNumber((_dafny.Seq.UnicodeFromString("icbmca")).length)),!(false)));
      }
    };
    static fm26(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(((!(false)) ? (_dafny.Set.fromElements(true, (_module.D13.create_DC31(false, !(true), new BigNumber(982))).dtor_cf48)) : (_dafny.Set.fromElements(false, true))), _dafny.Set.fromElements(true, false, !(!(!(!(false)))), true, !(true)), _dafny.Set.fromElements(false), _dafny.Set.fromElements(false));
    };
    static fm27(globalState) {
      return (_dafny.Set.fromElements(_module.D2.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-330)), function (_29_i0) {
  return _module.D1.create_DC2();
})))).Difference((_dafny.Set.fromElements(_module.D2.create_DC3(_dafny.Seq.of(_module.D1.create_DC2(), _module.D1.create_DC2())), _module.D2.create_DC3(_dafny.Seq.of(_module.D1.create_DC2(), _module.D1.create_DC2())))).Difference(function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of (_dafny.Set.fromElements(_module.D2.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_30_i1) {
  return _module.D1.create_DC2();
})), _module.D2.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), function (_31_i2) {
  return _module.D1.create_DC2();
})))).Elements) {
          let _32_v0 = _compr_7;
          if ((_dafny.Set.fromElements(_module.D2.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_30_i1) {
  return _module.D1.create_DC2();
})), _module.D2.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), function (_31_i2) {
  return _module.D1.create_DC2();
})))).contains(_32_v0)) {
            _coll7.add(_32_v0);
          }
        }
        return _coll7;
      }()));
    };
    static fm28(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(772),_dafny.areEqual(_dafny.Seq.UnicodeFromString("dwyanbw"), _dafny.Seq.UnicodeFromString("hg")));
    };
    static fm29(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-642)), function (_33_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length),new BigNumber(-52))).Merge(function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(129), new BigNumber(526))).cardinality()))).Elements) {
          let _34_v0 = _compr_8;
          if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(129), new BigNumber(526))).cardinality()))).contains(_34_v0)) {
            _coll8.push([(_34_v0).plus(new BigNumber(-126)),new BigNumber(390)]);
          }
        }
        return _coll8;
      }())).Merge(((true) ? (function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(958), new BigNumber(811))) {
          let _35_v1 = _compr_9;
          if (((new BigNumber(958)).isLessThanOrEqualTo(_35_v1)) && ((_35_v1).isLessThan(new BigNumber(811)))) {
            _coll9.push([_module.__default.safeModuloInt(_35_v1, new BigNumber(565)),new BigNumber(-456)]);
          }
        }
        return _coll9;
      }()) : (function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(516), new BigNumber(963))) {
          let _36_v2 = _compr_10;
          if (((new BigNumber(516)).isLessThanOrEqualTo(_36_v2)) && ((_36_v2).isLessThan(new BigNumber(963)))) {
            _coll10.push([(_36_v2).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(7))).length)),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length)]);
          }
        }
        return _coll10;
      }())));
    };
    static fm30(p0, p1, p2, p3, globalState) {
      if (true) {
        return _dafny.Seq.of(false, false);
      } else {
        return _dafny.Seq.of(!(true), !(false));
      }
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-589))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(584)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(true)).length)));
    };
    static m6(p0, p1, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _37_i0;
      _37_i0 = _dafny.ZERO;
      L0: {
        while (p0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_37_i0)) {
              break L0;
            }
            _37_i0 = (_37_i0).plus(_dafny.ONE);
            (globalState).f7 = true;
            let _38_v0;
            _38_v0 = new BigNumber(33);
            r1 = _38_v0;
            let _39_v1;
            let _nw0 = new _module.C1();
            _nw0.__ctor(false, p0, p0, _38_v0);
            _39_v1 = _nw0;
            _39_v1 = _39_v1;
            let _40_v2;
            let _nw1 = new _module.C0();
            _nw1.__ctor(p0, p0, (_39_v1).f19);
            _40_v2 = _nw1;
          }
        }
      }
      let _41_v3;
      let _nw2 = Array((new BigNumber(12)).toNumber()).fill(false);
      _41_v3 = _nw2;
      let _42_v4;
      _42_v4 = _dafny.Seq.of(_dafny.Seq.of(_41_v3, _41_v3));
      let _43_v5;
      _43_v5 = new BigNumber(98);
      let _44_v6;
      let _nw3 = Array((new BigNumber(3)).toNumber());
      _nw3[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.update((_42_v4)[_module.__default.safeIndex(_43_v5, new BigNumber((_42_v4).length))], _module.__default.safeIndex(_43_v5, new BigNumber(((_42_v4)[_module.__default.safeIndex(_43_v5, new BigNumber((_42_v4).length))]).length)), _41_v3)).length);
      _nw3[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_43_v5);
      _nw3[(new BigNumber(2)).toNumber()] = _43_v5;
      _44_v6 = _nw3;
      let _45_v7;
      _45_v7 = _module.D1.create_DC1(_44_v6);
      let _source2 = ((false) ? (_module.D1.create_DC1(_44_v6)) : (_45_v7));
      if (_source2.is_DC2) {
        let _46_v8;
        _46_v8 = new _dafny.CodePoint('r'.codePointAt(0));
        _46_v8 = _46_v8;
        let _47_v9;
        let _nw4 = new _module.C2();
        _nw4.__ctor(p0, _43_v5);
        _47_v9 = _nw4;
        _47_v9 = _47_v9;
        r2 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(170), _43_v5));
        let _48_v10;
        _48_v10 = _dafny.Map.Empty.slice().updateUnsafe(_43_v5,_43_v5);
        r1 = (((_48_v10).contains((_module.__default.fm0(globalState)).plus(new BigNumber(65)))) ? ((_48_v10).get((_module.__default.fm0(globalState)).plus(new BigNumber(65)))) : ((_dafny.ZERO).minus(_43_v5)));
      } else {
        let _49___mcc_h0 = (_source2).cf1;
        let _50_cf1 = _49___mcc_h0;
        let _51_v11;
        let _nw5 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
        _51_v11 = _nw5;
        let _52_v12;
        _52_v12 = _dafny.Map.Empty.slice().updateUnsafe(_43_v5,_43_v5);
        let _index0 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_51_v11).length));
        (_51_v11)[_index0] = _52_v12;
        let _53_v13;
        _53_v13 = _dafny.Seq.of(true, _module.__default.fm16(p0, _43_v5, p0, globalState));
        let _index1 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_51_v11).length));
        let _rhs0 = (_53_v13)[_module.__default.safeIndex(_43_v5, new BigNumber((_53_v13).length))];
        let _rhs1 = _43_v5;
        let _rhs2 = _module.__default.fm29(!(!(p0) || (p0)), globalState);
        let _lhs0 = globalState;
        let _lhs1 = _51_v11;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_51_v11).length));
        _lhs0.f7 = _rhs0;
        r0 = _rhs1;
        _lhs1[_lhs2] = _rhs2;
        (globalState).f7 = p0;
        (globalState).f7 = _module.__default.fm16(p0, new BigNumber(958), (p0) || (p0), globalState);
        let _index2 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_41_v3).length));
        (_41_v3)[_index2] = p0;
        let _index3 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_41_v3).length));
        (_41_v3)[_index3] = p0;
      }
      let _54_i1;
      _54_i1 = _dafny.ZERO;
      L1: {
        while (p0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_54_i1)) {
              break L1;
            }
            _54_i1 = (_54_i1).plus(_dafny.ONE);
            let _55_v14;
            let _nw6 = new _module.C2();
            _nw6.__ctor(p0, _43_v5);
            _55_v14 = _nw6;
            let _56_v15;
            _56_v15 = _dafny.Seq.of(true, p0, p0, p0);
            let _57_v16;
            _57_v16 = _dafny.Map.Empty.slice().updateUnsafe((_56_v15)[_module.__default.safeIndex(_43_v5, new BigNumber((_56_v15).length))],false);
            let _58_v17;
            let _nw7 = new _module.C1();
            _nw7.__ctor(p0, true, p0, new BigNumber((_57_v16).length));
            _58_v17 = _nw7;
            let _59_v18;
            _59_v18 = _dafny.Set.fromElements(_58_v17, _58_v17);
            if ((_dafny.Set.fromElements(_58_v17)).IsDisjointFrom(_59_v18)) {
              let _60_v19;
              _60_v19 = _module.D2.create_DC4();
              (_58_v17).f22 = (_58_v17).fm7(_60_v19, p1, globalState);
              let _61_v20;
              _61_v20 = _dafny.Seq.of(new BigNumber((_56_v15).length), _43_v5);
              let _62_v21;
              let _nw8 = new _module.C2();
              _nw8.__ctor(_58_v17.f22, (new BigNumber((_56_v15).length)).plus(new BigNumber((_61_v20).length)));
              _62_v21 = _nw8;
              let _63_v22;
              _63_v22 = new _dafny.CodePoint('j'.codePointAt(0));
              let _64_v23;
              let _init0 = ((_65_v15, _66_v5) => function (_67_i2) {
                return _dafny.Seq.of(_65_v15, _dafny.Seq.update(_65_v15, _module.__default.safeIndex(_66_v5, new BigNumber((_65_v15).length)), false));
              })(_56_v15, _43_v5);
              let _nw9 = Array((new BigNumber(2)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw9.length); _i0_0++) {
                _nw9[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _64_v23 = _nw9;
              let _68_v24;
              _68_v24 = _dafny.Seq.of(_56_v15);
              let _index4 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_64_v23).length));
              (_64_v23)[_index4] = _68_v24;
              let _69_v25;
              _69_v25 = _dafny.Seq.UnicodeFromString("kjypqixxw");
              let _70_v26;
              let _nw10 = new _module.C3();
              _nw10.__ctor(_41_v3, _69_v25, (_58_v17).f23, _43_v5);
              _70_v26 = _nw10;
              let _71_v27;
              _71_v27 = _dafny.Map.Empty.slice().updateUnsafe((_58_v17).f23,_43_v5);
              let _72_v28;
              _72_v28 = _dafny.Set.fromElements(_58_v17.f22, p0);
              let _73_v29;
              _73_v29 = _module.D3.create_DC7(_70_v26, _71_v27, p0, _72_v28, _58_v17.f22);
              let _index5 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_64_v23).length));
              let _rhs3 = _58_v17.f22;
              let _rhs4 = _62_v21;
              let _rhs5 = (_55_v14).fm4(_58_v17.f22, _dafny.Seq.contains(_69_v25, _63_v22), (_73_v29).dtor_cf6, globalState);
              let _rhs6 = _58_v17.f22;
              let _rhs7 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), ((_74_v15) => function (_75_i3) {
                return _74_v15;
              })(_56_v15)), _68_v24);
              let _lhs3 = globalState;
              let _lhs4 = _58_v17;
              let _lhs5 = _64_v23;
              let _lhs6 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_64_v23).length));
              _lhs3.f7 = _rhs3;
              _62_v21 = _rhs4;
              _63_v22 = _rhs5;
              _lhs4.f22 = _rhs6;
              _lhs5[_lhs6] = _rhs7;
              let _76_v30;
              _76_v30 = _dafny.Set.fromElements(p1);
              let _77_v31;
              let _nw11 = new _module.C3();
              _nw11.__ctor(_41_v3, _69_v25, _58_v17.f22, new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(_63_v22, _63_v22, p1, new _dafny.CodePoint('q'.codePointAt(0))), _76_v30, _76_v30)).length));
              _77_v31 = _nw11;
              let _78_v32;
              _78_v32 = _dafny.MultiSet.fromElements(_77_v31, _77_v31);
              let _79_v33;
              _79_v33 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm30(_43_v5, (_70_v26).f19, _43_v5, _43_v5, globalState)),_78_v32);
              let _80_v34;
              _80_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,_79_v33);
              let _81_v35;
              _81_v35 = _dafny.Seq.of((((_80_v34).contains((_58_v17).f23)) ? ((_80_v34).get((_58_v17).f23)) : ((_79_v33).update(_56_v15, _78_v32))));
              _79_v33 = (_81_v35)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_69_v25, _dafny.Seq.UnicodeFromString("rlfd")), _module.__default.safeIndex(_43_v5, new BigNumber((_dafny.Seq.Concat(_69_v25, _dafny.Seq.UnicodeFromString("rlfd"))).length)), _63_v22)).length), new BigNumber((_81_v35).length))];
              let _index6 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_44_v6).length));
              (_44_v6)[_index6] = new BigNumber(324);
              let _index7 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_44_v6).length));
              (_44_v6)[_index7] = (((_56_v15)[_module.__default.safeIndex((_70_v26).f19, new BigNumber((_56_v15).length))]) ? (((_70_v26).f19).plus(_43_v5)) : (_43_v5));
              let _82_v36;
              _82_v36 = _module.D10.create_DC24((new BigNumber(-233)).plus(_43_v5), _module.__default.safeDivisionInt((_44_v6)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_44_v6).length))], (_44_v6)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_44_v6).length))]), new BigNumber((_71_v27).length), p0, (_77_v31).fm2(_69_v25, new BigNumber(-907), globalState));
              _82_v36 = _module.D10.create_DC24((_module.__default.fm0(globalState)).plus(new BigNumber(283)), (_44_v6)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_44_v6).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bs"), _dafny.Seq.UnicodeFromString("macylwhcb"))).length), (_70_v26).f18, (_58_v17).f23);
            } else {
              _57_v16 = (_57_v16).update(p0, _module.__default.fm16(_58_v17.f22, _43_v5, !(true), globalState));
              let _83_v37;
              _83_v37 = _dafny.MultiSet.fromElements((_58_v17).f23, p0, _58_v17.f22, (_58_v17).f23);
              let _rhs8 = !(false);
              let _rhs9 = !(((_dafny.MultiSet.fromElements((_58_v17).f23, _58_v17.f22)).Intersect(_dafny.MultiSet.FromArray(_56_v15))).IsProperSubsetOf(_83_v37));
              let _lhs7 = globalState;
              let _lhs8 = _58_v17;
              _lhs7.f7 = _rhs8;
              _lhs8.f22 = _rhs9;
              r2 = _43_v5;
              let _84_v38;
              _84_v38 = _dafny.Seq.of(_43_v5);
              let _85_v39;
              _85_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_84_v38).length),p0);
              let _86_v40;
              _86_v40 = _dafny.Map.Empty.slice().updateUnsafe((((_85_v39).contains(_43_v5)) ? ((_85_v39).get(_43_v5)) : ((_58_v17).f23)),(_55_v14).fm5(_58_v17.f22, p1, (_58_v17).f23, (_58_v17).f23, globalState));
              let _87_v41;
              _87_v41 = _module.D2.create_DC4();
              _86_v40 = (_86_v40).update((_58_v17).fm7(_87_v41, p1, globalState), _83_v37);
              (_58_v17).f22 = !((_58_v17.f22) && (_58_v17.f22));
            }
            let _88_v42;
            let _89_v43;
            let _90_v44;
            let _91_v45;
            let _out0;
            let _out1;
            let _out2;
            let _out3;
            let _outcollector0 = (_55_v14).m2(p0, _58_v17.f22, p0, _module.__default.fm0(globalState), globalState);
            _out0 = _outcollector0[0];
            _out1 = _outcollector0[1];
            _out2 = _outcollector0[2];
            _out3 = _outcollector0[3];
            _88_v42 = _out0;
            _89_v43 = _out1;
            _90_v44 = _out2;
            _91_v45 = _out3;
            let _92_v46;
            _92_v46 = _dafny.Seq.of(_44_v6, _44_v6, _44_v6);
            let _93_v47;
            _93_v47 = _module.D6.create_DC12(_dafny.Seq.update(_dafny.Seq.Concat(_92_v46, _92_v46), _module.__default.safeIndex(_43_v5, new BigNumber((_dafny.Seq.Concat(_92_v46, _92_v46)).length)), _44_v6));
            let _source3 = _93_v47;
            if (_source3.is_DC13) {
              let _94___mcc_h1 = (_source3).cf18;
              let _95_cf18 = _94___mcc_h1;
              let _96_v48;
              _96_v48 = _dafny.MultiSet.fromElements(_91_v45);
              let _97_v49;
              _97_v49 = _dafny.Seq.of(_96_v48);
              let _98_v50;
              _98_v50 = _module.D4.create_DC8(new BigNumber((_97_v49).length));
              let _99_v51;
              _99_v51 = _dafny.Map.Empty.slice().updateUnsafe(_98_v50,_55_v14);
              _99_v51 = (_99_v51).Merge(_99_v51);
              let _index8 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_41_v3).length));
              (_41_v3)[_index8] = _58_v17.f22;
              let _100_v52;
              _100_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,_41_v3);
              let _index9 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_41_v3).length));
              (_41_v3)[_index9] = !((_100_v52).contains((_58_v17).f23));
              r2 = _91_v45;
              _57_v16 = (_57_v16).update(((_dafny.ZERO).minus(_95_cf18)).isLessThan(_95_cf18), p0);
            } else {
              let _101___mcc_h2 = (_source3).cf17;
              let _102_cf17 = _101___mcc_h2;
              let _103_v53;
              let _nw12 = Array((new BigNumber(13)).toNumber());
              _103_v53 = _nw12;
              let _104_v54;
              let _nw13 = new _module.C0();
              _nw13.__ctor(_58_v17.f22, p0, _91_v45);
              _104_v54 = _nw13;
              let _index10 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_103_v53).length));
              (_103_v53)[_index10] = _104_v54;
              let _index11 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_103_v53).length));
              (_103_v53)[_index11] = _104_v54;
              let _index12 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_44_v6).length));
              (_44_v6)[_index12] = _91_v45;
              let _index13 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_44_v6).length));
              (_44_v6)[_index13] = (_dafny.ZERO).minus((_91_v45).plus(_91_v45));
              _91_v45 = _43_v5;
              (globalState).f7 = !(_89_v43);
            }
          }
        }
      }
      let _105_v55;
      _105_v55 = _dafny.Seq.of(_41_v3);
      let _106_v56;
      _106_v56 = _dafny.MultiSet.fromElements((new BigNumber(711)).minus(new BigNumber(-106)), _module.__default.safeModuloInt(_43_v5, new BigNumber((_105_v55).length)), _43_v5);
      let _index14 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length));
      (_41_v3)[_index14] = p0;
      let _107_v57;
      let _init1 = ((_108_p0) => function (_109_i4) {
        return _dafny.Set.fromElements(_108_p0, _108_p0);
      })(p0);
      let _nw14 = Array((new BigNumber(13)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw14.length); _i0_1++) {
        _nw14[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _107_v57 = _nw14;
      let _110_v58;
      _110_v58 = _dafny.Set.fromElements(p0, p0, p0);
      let _index15 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_107_v57).length));
      (_107_v57)[_index15] = (_110_v58).Union(_110_v58);
      let _111_v59;
      _111_v59 = _dafny.Seq.UnicodeFromString("bcfyxq");
      let _112_v60;
      let _nw15 = new _module.C3();
      _nw15.__ctor(_41_v3, _111_v59, p0, _43_v5);
      _112_v60 = _nw15;
      let _113_v61;
      _113_v61 = _module.D10.create_DC23(p1);
      let _114_v62;
      _114_v62 = _module.D10.create_DC25(_113_v61);
      let _115_v63;
      _115_v63 = _module.D3.create_DC7(_112_v60, _module.__default.fm31(_114_v62, p0, (_112_v60).f19, (_112_v60).f19, globalState), (_112_v60).f18, _110_v58, p0);
      let _116_v64;
      _116_v64 = _dafny.Set.fromElements(_112_v60);
      let _117_v65;
      _117_v65 = _dafny.Set.fromElements(_44_v6, _44_v6);
      let _118_v66;
      _118_v66 = _module.D8.create_DC19(_41_v3, p0, _116_v64, _117_v65, p0);
      let _119_v67;
      let _nw16 = new _module.C0();
      _nw16.__ctor(p0, (_118_v66).dtor_cf25, _43_v5);
      _119_v67 = _nw16;
      let _120_v68;
      _120_v68 = _dafny.Map.Empty.slice().updateUnsafe(_115_v63,_119_v67);
      let _121_v69;
      _121_v69 = _dafny.Set.fromElements(_120_v68);
      let _index16 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length));
      let _index17 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_107_v57).length));
      let _rhs10 = _dafny.MultiSet.fromElements(_43_v5, _43_v5, _43_v5, _43_v5);
      let _rhs11 = (_121_v69).IsSubsetOf(_121_v69);
      let _rhs12 = _110_v58;
      let _lhs9 = _41_v3;
      let _lhs10 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length));
      let _lhs11 = _107_v57;
      let _lhs12 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_107_v57).length));
      _106_v56 = _rhs10;
      _lhs9[_lhs10] = _rhs11;
      _lhs11[_lhs12] = _rhs12;
      let _122_v70;
      _122_v70 = _dafny.Seq.of(_43_v5, (_112_v60).f19);
      let _123_v71;
      _123_v71 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_122_v70).length),new BigNumber(906));
      let _124_v72;
      _124_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_123_v71).length),(_119_v67).f24);
      let _125_v73;
      _125_v73 = _dafny.MultiSet.fromElements((((_124_v72).contains(_43_v5)) ? ((_124_v72).get(_43_v5)) : (p0)));
      let _index18 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length));
      (_41_v3)[_index18] = ((_dafny.Set.fromElements((_119_v67).f24)).Difference((_107_v57)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_107_v57).length))])).IsProperSubsetOf((_module.__default.fm10(_43_v5, new BigNumber(((_125_v73)).cardinality()), globalState)).Difference(_110_v58));
      let _126_v74;
      _126_v74 = _module.D6.create_DC12(_dafny.Seq.of(_44_v6, _44_v6, _44_v6));
      let _source4 = _126_v74;
      if (_source4.is_DC13) {
        let _127___mcc_h3 = (_source4).cf18;
        let _128_cf18 = _127___mcc_h3;
        let _129_v75;
        _129_v75 = new _dafny.CodePoint('o'.codePointAt(0));
        _129_v75 = p1;
        (globalState).f7 = _module.__default.fm16((_41_v3)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length))], new BigNumber(((((_112_v60).f18) ? (_dafny.Seq.UnicodeFromString("axwddu")) : (_dafny.Seq.UnicodeFromString("irxoor")))).length), (_119_v67).f24, globalState);
        let _130_v76;
        let _nw17 = new _module.C2();
        _nw17.__ctor((_112_v60).f18, _module.__default.fm0(globalState));
        _130_v76 = _nw17;
        let _rhs13 = (_dafny.ZERO).minus(((_128_cf18).multipliedBy(_128_cf18)).multipliedBy((_112_v60).f19));
        let _rhs14 = ((p0) ? (_130_v76) : ((((_119_v67).f24) ? (_130_v76) : (_130_v76))));
        r1 = _rhs13;
        _130_v76 = _rhs14;
        let _131_v77;
        let _nw18 = new _module.C2();
        _nw18.__ctor((_119_v67).f24, (_112_v60).f19);
        _131_v77 = _nw18;
      } else {
        let _132___mcc_h4 = (_source4).cf17;
        let _133_cf17 = _132___mcc_h4;
        let _134_v78;
        _134_v78 = _dafny.Seq.of((_41_v3)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length))]);
        _134_v78 = _dafny.Seq.Concat(_dafny.Seq.of((_41_v3)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length))], (_119_v67).f24, (_119_v67).f24, (_112_v60).f18, (_41_v3)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length))]), _134_v78);
        let _135_v79;
        _135_v79 = _module.D6.create_DC13((_112_v60).f19);
        let _source5 = _135_v79;
        if (_source5.is_DC13) {
          let _136___mcc_h5 = (_source5).cf18;
          let _137_cf18 = _136___mcc_h5;
          let _138_v80;
          _138_v80 = _dafny.Seq.of(_110_v58, _110_v58, _module.__default.fm10((_112_v60).f19, new BigNumber((_122_v70).length), globalState));
          r2 = new BigNumber((_138_v80).length);
          let _139_v81;
          _139_v81 = _module.D4.create_DC8(_43_v5);
          let _140_v82;
          _140_v82 = _dafny.Map.Empty.slice().updateUnsafe(((_112_v60).f19).isEqualTo(_137_cf18),((_139_v81).dtor_cf10).isLessThan((_112_v60).f19));
          let _index19 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length));
          (_41_v3)[_index19] = (((_140_v82).contains((_119_v67).f24)) ? ((_140_v82).get((_119_v67).f24)) : (true));
          (globalState).f7 = (((_124_v72).contains((_112_v60).f19)) ? ((_124_v72).get((_112_v60).f19)) : (!((_41_v3)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length))])));
          let _141_v83;
          let _nw19 = Array((new BigNumber(28)).toNumber());
          _141_v83 = _nw19;
          let _142_v84;
          let _nw20 = new _module.C2();
          _nw20.__ctor((_112_v60).f18, (_112_v60).f19);
          _142_v84 = _nw20;
          let _index20 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_141_v83).length));
          (_141_v83)[_index20] = _142_v84;
          let _143_v85;
          let _nw21 = Array((_dafny.ONE).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = _112_v60;
          _143_v85 = _nw21;
          let _index21 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_143_v85).length));
          (_143_v85)[_index21] = _112_v60;
          let _index22 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_141_v83).length));
          let _index23 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_143_v85).length));
          let _rhs15 = _142_v84;
          let _rhs16 = _119_v67;
          let _rhs17 = !(((new BigNumber(345)).multipliedBy((_112_v60).f19)).isLessThanOrEqualTo(new BigNumber((_111_v59).length)));
          let _rhs18 = _112_v60;
          let _rhs19 = _111_v59;
          let _lhs13 = _141_v83;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_141_v83).length));
          let _lhs15 = globalState;
          let _lhs16 = _143_v85;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_143_v85).length));
          _lhs13[_lhs14] = _rhs15;
          _119_v67 = _rhs16;
          _lhs15.f7 = _rhs17;
          _lhs16[_lhs17] = _rhs18;
          _111_v59 = _rhs19;
        } else {
          let _144___mcc_h6 = (_source5).cf17;
          let _145_cf17 = _144___mcc_h6;
          (globalState).f7 = (_134_v78)[_module.__default.safeIndex((((_119_v67).f24) ? (new BigNumber(899)) : (_module.__default.fm0(globalState))), new BigNumber((_134_v78).length))];
          _111_v59 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(434)), ((_146_p1) => function (_147_i5) {
            return _146_p1;
          })(p1));
          let _148_v86;
          _148_v86 = _module.D5.create_DC11((_112_v60).f18, new BigNumber((_111_v59).length), (_112_v60).f18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(893)), ((_149_p1) => function (_150_i6) {
  return _149_p1;
})(p1)), (_112_v60).f19);
          (globalState).f7 = _dafny.areEqual((_148_v86).dtor_cf15, _111_v59);
          r1 = (_112_v60).f19;
        }
        let _151_v87;
        let _nw22 = new _module.C2();
        _nw22.__ctor(false, (_112_v60).f19);
        _151_v87 = _nw22;
        if (true) {
          let _152_v88;
          _152_v88 = _dafny.Map.Empty.slice().updateUnsafe((_45_v7).dtor_cf1,p1);
          _152_v88 = (_152_v88).update(_44_v6, p1);
          _110_v58 = _110_v58;
          let _153_v89;
          _153_v89 = _dafny.Map.Empty.slice().updateUnsafe(!((_43_v5).isEqualTo((_dafny.ZERO).minus((_112_v60).f19))),(_module.__default.fm16((_112_v60).f18, (_112_v60).f19, (_119_v67).f24, globalState)) || ((((_124_v72).contains((_112_v60).f19)) ? ((_124_v72).get((_112_v60).f19)) : ((_41_v3)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_41_v3).length))]))));
          let _154_v90;
          _154_v90 = _module.D5.create_DC11((_119_v67).f24, (_112_v60).f19, true, _dafny.Seq.UnicodeFromString("nhjboacjp"), (_112_v60).f19);
          let _155_v91;
          _155_v91 = _dafny.Set.fromElements((_112_v60).f19, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(560)), ((_156_p1) => function (_157_i7) {
            return _156_p1;
          })(p1))).length), (_122_v70)[_module.__default.safeIndex((_154_v90).dtor_cf13, new BigNumber((_122_v70).length))], (_112_v60).f19, (((_123_v71).contains((_112_v60).f19)) ? ((_123_v71).get((_112_v60).f19)) : ((_112_v60).f19)));
          _153_v89 = (_153_v89).update((_155_v91).IsSubsetOf(_155_v91), (_119_v67).f24);
          (globalState).f7 = !(!(_module.__default.fm16((((_124_v72).contains((_112_v60).f19)) ? ((_124_v72).get((_112_v60).f19)) : (false)), (_112_v60).f19, (_112_v60).f18, globalState)));
          (globalState).f7 = ((_107_v57)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_107_v57).length))]).IsSubsetOf((_107_v57)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_107_v57).length))]);
        } else {
          let _158_v92;
          let _nw23 = new _module.C0();
          _nw23.__ctor((_dafny.Set.fromElements((_112_v60).f18)).IsSubsetOf(_110_v58), (_112_v60).f18, new BigNumber(-352));
          _158_v92 = _nw23;
          _111_v59 = _111_v59;
          let _159_v93;
          _159_v93 = _module.D11.create_DC27((_112_v60).f19, (_112_v60).f19, (_112_v60).f19, _106_v56, _122_v70);
          r0 = (_159_v93).dtor_cf40;
          _111_v59 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(102)), ((_160_p1) => function (_161_i8) {
            return _160_p1;
          })(p1));
          r1 = (_dafny.ZERO).minus((_112_v60).f19);
        }
      }
      r0 = _43_v5;
      let _162_v94;
      _162_v94 = _module.D6.create_DC13((_112_v60).f19);
      let _pat_let_tv0 = _119_v67;
      let _pat_let_tv1 = _112_v60;
      let _pat_let_tv2 = _122_v70;
      let _pat_let_tv3 = _112_v60;
      let _pat_let_tv4 = _112_v60;
      let _pat_let_tv5 = _119_v67;
      let _pat_let_tv6 = _43_v5;
      let _pat_let_tv7 = _111_v59;
      let _pat_let_tv8 = _106_v56;
      let _pat_let_tv9 = _111_v59;
      let _pat_let_tv10 = p1;
      r1 = function (_source6) {
        if (_source6.is_DC13) {
          let _163___mcc_h7 = (_source6).cf18;
          let _164_cf18 = _163___mcc_h7;
          if ((_pat_let_tv0).f24) {
            return (_module.D10.create_DC24((_pat_let_tv1).f19, new BigNumber((_pat_let_tv2).length), (_pat_let_tv3).f19, (_pat_let_tv4).f18, (_pat_let_tv5).f24)).dtor_cf34;
          } else {
            return _pat_let_tv6;
          }
        } else {
          let _165___mcc_h8 = (_source6).cf17;
          let _166_cf17 = _165___mcc_h8;
          return new BigNumber((_dafny.Seq.update(_pat_let_tv7, _module.__default.safeIndex(new BigNumber((_pat_let_tv8).cardinality()), new BigNumber((_pat_let_tv9).length)), _pat_let_tv10)).length);
        }
      }(_162_v94);
      let _167_v95;
      _167_v95 = _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), function (_168_i9) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _module.__default.safeIndex((_122_v70)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_111_v59).length)), new BigNumber((_122_v70).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), function (_169_i9) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      })).length)), p1), _module.__default.safeIndex(_43_v5, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), function (_170_i9) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _module.__default.safeIndex((_122_v70)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_111_v59).length)), new BigNumber((_122_v70).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), function (_171_i9) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      })).length)), p1)).length)), p1), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-676)), ((_172_p1) => function (_173_i10) {
        return _172_p1;
      })(p1)));
      r2 = new BigNumber(((_167_v95).Intersect((_167_v95).Union(_dafny.Set.fromElements(_111_v59, _111_v59)))).length);
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _174_v0;
      _174_v0 = false;
      let _175_v1;
      _175_v1 = _dafny.Seq.of(!(_174_v0));
      let _176_v2;
      _176_v2 = new BigNumber(173);
      let _177_v3;
      let _nw24 = Array((new BigNumber(28)).toNumber());
      _nw24[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_174_v0, !(_174_v0));
      _nw24[(_dafny.ONE).toNumber()] = _175_v1;
      _nw24[(new BigNumber(2)).toNumber()] = (_175_v1);
      _nw24[(new BigNumber(3)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(4)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(5)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_175_v1, _module.__default.safeIndex(new BigNumber(-540), new BigNumber((_175_v1).length)), _174_v0);
      _nw24[(new BigNumber(7)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_174_v0, _174_v0);
      _nw24[(new BigNumber(9)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(_174_v0, _174_v0);
      _nw24[(new BigNumber(11)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_175_v1, _module.__default.safeIndex(_176_v2, new BigNumber((_175_v1).length)), _174_v0);
      _nw24[(new BigNumber(13)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(_174_v0);
      _nw24[(new BigNumber(15)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(16)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(_174_v0, _174_v0);
      _nw24[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(_174_v0, _174_v0);
      _nw24[(new BigNumber(19)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(20)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(21)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(22)).toNumber()] = _dafny.Seq.of(_174_v0);
      _nw24[(new BigNumber(23)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(24)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(25)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(26)).toNumber()] = _175_v1;
      _nw24[(new BigNumber(27)).toNumber()] = _175_v1;
      _177_v3 = _nw24;
      let _178_v4;
      _178_v4 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_174_v0);
      let _179_v5;
      _179_v5 = _175_v1;
      let _180_v6;
      let _nw25 = Array((new BigNumber(10)).toNumber());
      _nw25[(_dafny.ZERO).toNumber()] = new BigNumber(963);
      _nw25[(_dafny.ONE).toNumber()] = _176_v2;
      _nw25[(new BigNumber(2)).toNumber()] = _176_v2;
      _nw25[(new BigNumber(3)).toNumber()] = new BigNumber((_178_v4).length);
      _nw25[(new BigNumber(4)).toNumber()] = _176_v2;
      _nw25[(new BigNumber(5)).toNumber()] = _176_v2;
      _nw25[(new BigNumber(6)).toNumber()] = new BigNumber(782);
      _nw25[(new BigNumber(7)).toNumber()] = _176_v2;
      _nw25[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_175_v1,_176_v2)).length);
      _nw25[(new BigNumber(9)).toNumber()] = _176_v2;
      _180_v6 = _nw25;
      let _181_v7;
      _181_v7 = _module.D1.create_DC1(_180_v6);
      let _182_v8;
      let _nw26 = Array((new BigNumber(28)).toNumber());
      _nw26[(_dafny.ZERO).toNumber()] = _180_v6;
      _nw26[(_dafny.ONE).toNumber()] = _180_v6;
      _nw26[(new BigNumber(2)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(3)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(4)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(5)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(6)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(7)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(8)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(9)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(10)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(11)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(12)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(13)).toNumber()] = (_181_v7).dtor_cf1;
      _nw26[(new BigNumber(14)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(15)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(16)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(17)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(18)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(19)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(20)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(21)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(22)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(23)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(24)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(25)).toNumber()] = _180_v6;
      _nw26[(new BigNumber(26)).toNumber()] = (_181_v7).dtor_cf1;
      _nw26[(new BigNumber(27)).toNumber()] = _180_v6;
      _182_v8 = _nw26;
      let _183_v9;
      _183_v9 = new _dafny.CodePoint('o'.codePointAt(0));
      let _184_v10;
      _184_v10 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_183_v9);
      let _185_v11;
      _185_v11 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,(((_184_v10).contains(_176_v2)) ? ((_184_v10).get(_176_v2)) : (_183_v9)));
      let _186_v12;
      let _nw27 = Array((new BigNumber(10)).toNumber());
      _nw27[(_dafny.ZERO).toNumber()] = _174_v0;
      _nw27[(_dafny.ONE).toNumber()] = true;
      _nw27[(new BigNumber(2)).toNumber()] = _174_v0;
      _nw27[(new BigNumber(3)).toNumber()] = _174_v0;
      _nw27[(new BigNumber(4)).toNumber()] = _174_v0;
      _nw27[(new BigNumber(5)).toNumber()] = _174_v0;
      _nw27[(new BigNumber(6)).toNumber()] = _174_v0;
      _nw27[(new BigNumber(7)).toNumber()] = _174_v0;
      _nw27[(new BigNumber(8)).toNumber()] = _174_v0;
      _nw27[(new BigNumber(9)).toNumber()] = _174_v0;
      _186_v12 = _nw27;
      let _187_v13;
      _187_v13 = _dafny.Seq.of(_186_v12, _186_v12);
      let _188_v15;
      _188_v15 = _dafny.MultiSet.fromElements(_184_v10, _185_v11, (_185_v11).update(new BigNumber((_187_v13).length), _183_v9), function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of (_184_v10).Keys.Elements) {
          let _189_v14 = _compr_11;
          if ((_184_v10).contains(_189_v14)) {
            _coll11.push([_module.__default.safeModuloInt(_189_v14, _176_v2),_183_v9]);
          }
        }
        return _coll11;
      }(), _185_v11);
      let _190_v16;
      _190_v16 = _dafny.Map.Empty.slice().updateUnsafe(_174_v0,_174_v0);
      let _191_v17;
      _191_v17 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,new BigNumber((_178_v4).length));
      let _192_v18;
      _192_v18 = _dafny.Seq.UnicodeFromString("hf");
      let _193_v19;
      _193_v19 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ldynryh"), _192_v18);
      let _194_v20;
      _194_v20 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_186_v12);
      let _195_globalState;
      let _nw28 = new _module.GlobalState();
      _nw28.__ctor(_177_v3, _182_v8, new BigNumber(922), _188_v15, new BigNumber(944), new _dafny.CodePoint('l'.codePointAt(0)), (_dafny.Map.Empty.slice().updateUnsafe(_174_v0,_174_v0)).Merge((_190_v16).update(_174_v0, _174_v0)), true, new BigNumber(-334), _191_v17, true, new BigNumber(303), new BigNumber(294), false, false, _193_v19, _194_v20, false);
      _195_globalState = _nw28;
      let _hi0 = new BigNumber(-973);
      for (let _196_i0 = _module.__default.fm0(_195_globalState); _196_i0.isLessThan(_hi0); _196_i0 = _196_i0.plus(_dafny.ONE)) {
        let _197_v21;
        _197_v21 = _dafny.Set.fromElements(_176_v2, _196_i0, _176_v2);
        let _198_v22;
        _198_v22 = _dafny.Map.Empty.slice().updateUnsafe(_174_v0,new BigNumber((_197_v21).length));
        let _rhs20 = _module.__default.safeModuloInt(_176_v2, _176_v2);
        let _rhs21 = new BigNumber((_192_v18).length);
        let _rhs22 = (_175_v1)[_module.__default.safeIndex(new BigNumber(((_198_v22).Merge(_198_v22)).length), new BigNumber((_175_v1).length))];
        _176_v2 = _rhs20;
        _176_v2 = _rhs21;
        _174_v0 = _rhs22;
        _176_v2 = _196_i0;
        _176_v2 = (((_174_v0) && (_174_v0)) ? (new BigNumber(173)) : (_176_v2));
        let _199_v23;
        let _nw29 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _199_v23 = _nw29;
        _199_v23 = _199_v23;
      }
      let _index24 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
      (_186_v12)[_index24] = _174_v0;
      let _index25 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
      (_186_v12)[_index25] = _dafny.Seq.IsProperPrefixOf(_192_v18, _192_v18);
      _176_v2 = _176_v2;
      let _200_v24;
      _200_v24 = _dafny.Map.Empty.slice().updateUnsafe((_175_v1)[_module.__default.safeIndex(_176_v2, new BigNumber((_175_v1).length))],_186_v12);
      let _rhs23 = _174_v0;
      let _rhs24 = ((_200_v24).Merge(_200_v24)).Merge(_200_v24);
      let _lhs18 = _195_globalState;
      _lhs18.f7 = _rhs23;
      _200_v24 = _rhs24;
      if ((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))]) {
        let _nw30 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _180_v6 = _nw30;
        _176_v2 = _176_v2;
        let _201_v25;
        _201_v25 = _dafny.Map.Empty.slice().updateUnsafe(_183_v9,_176_v2);
        let _202_v26;
        let _nw31 = new _module.C2();
        _nw31.__ctor((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))], (((_201_v25).contains(_183_v9)) ? ((_201_v25).get(_183_v9)) : (new BigNumber(43))));
        _202_v26 = _nw31;
        let _203_v27;
        _203_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-888)), ((_204_v9) => function (_205_i1) {
          return _204_v9;
        })(_183_v9)),_176_v2);
        _203_v27 = (_203_v27).update(_module.__default.fm9(new BigNumber(835), _195_globalState), new BigNumber(-490));
        let _206_v28;
        _206_v28 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm26(_175_v1, (_dafny.ZERO).minus((((_201_v25).contains(_183_v9)) ? ((_201_v25).get(_183_v9)) : (_176_v2))), _195_globalState),(_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))]);
        let _207_v29;
        _207_v29 = _dafny.Set.fromElements(_174_v0, (_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))], (_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))]);
        let _208_v30;
        let _nw32 = new _module.C3();
        _nw32.__ctor(_186_v12, _dafny.Seq.UnicodeFromString("ssio"), (_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))], _176_v2);
        _208_v30 = _nw32;
        let _209_v31;
        _209_v31 = _dafny.Map.Empty.slice().updateUnsafe((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))],_176_v2);
        let _210_v32;
        _210_v32 = _module.D3.create_DC7(_208_v30, _209_v31, (_208_v30).f18, _207_v29, true);
        let _211_v33;
        _211_v33 = _dafny.MultiSet.fromElements(_207_v29, _207_v29, _207_v29, _207_v29, (_210_v32).dtor_cf8);
        _206_v28 = (_206_v28).update(_211_v33, !((_178_v4).equals(_178_v4)));
      } else {
        let _212_v34;
        _212_v34 = _module.D1.create_DC2();
        let _index26 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
        let _rhs25 = _212_v34;
        let _rhs26 = (_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))];
        let _lhs19 = _186_v12;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
        _212_v34 = _rhs25;
        _lhs19[_lhs20] = _rhs26;
        let _213_v35;
        _213_v35 = _dafny.MultiSet.fromElements(_174_v0);
        let _214_v36;
        _214_v36 = _dafny.Map.Empty.slice().updateUnsafe(_183_v9,(_dafny.ZERO).minus(new BigNumber((_213_v35).cardinality())));
        let _215_v37;
        _215_v37 = _module.D9.create_DC20(_214_v36);
        let _pat_let_tv11 = _214_v36;
        _214_v36 = (function (_pat_let0_0) {
          return function (_216_dt__update__tmp_h1) {
            return function (_pat_let1_0) {
              return function (_217_dt__update_hcf29_h0) {
                return _module.D9.create_DC20(_217_dt__update_hcf29_h0);
              }(_pat_let1_0);
            }(_pat_let_tv11);
          }(_pat_let0_0);
        }(_215_v37)).dtor_cf29;
        let _218_v38;
        _218_v38 = _module.D10.create_DC23(_183_v9);
        let _219_v39;
        let _220_v40;
        let _221_v41;
        let _out4;
        let _out5;
        let _out6;
        let _outcollector1 = _module.__default.m6((_176_v2).isLessThan(_176_v2), (_218_v38).dtor_cf31, _195_globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _out6 = _outcollector1[2];
        _219_v39 = _out4;
        _220_v40 = _out5;
        _221_v41 = _out6;
        let _222_v42;
        _222_v42 = _dafny.Seq.of(_212_v34);
        let _223_v43;
        _223_v43 = _module.D2.create_DC3(_222_v42);
        let _224_v44;
        _224_v44 = _dafny.MultiSet.fromElements(_module.D2.create_DC3(_dafny.Seq.of(_212_v34)), _223_v43);
        let _225_v46;
        _225_v46 = _dafny.Set.fromElements(_module.D2.create_DC3(_222_v42));
        let _226_v47;
        _226_v47 = _dafny.Seq.of(function () {
          let _coll12 = new _dafny.Set();
          for (const _compr_12 of (_224_v44).Elements) {
            let _227_v45 = _compr_12;
            if ((_224_v44).contains(_227_v45)) {
              _coll12.add(_227_v45);
            }
          }
          return _coll12;
        }(), _225_v46);
        (_195_globalState).f7 = (((_226_v47)[_module.__default.safeIndex(_220_v40, new BigNumber((_226_v47).length))]).Union(_225_v46)).IsDisjointFrom(_module.__default.fm27(_195_globalState));
        let _228_v48;
        let _nw33 = new _module.C1();
        _nw33.__ctor((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))], _174_v0, _174_v0, _220_v40);
        _228_v48 = _nw33;
        let _229_v49;
        _229_v49 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_228_v48);
        let _230_v50;
        _230_v50 = _dafny.Set.fromElements(_229_v49);
        let _231_v51;
        _231_v51 = _dafny.MultiSet.fromElements(_219_v39, _221_v41);
        let _232_v52;
        _232_v52 = _dafny.Seq.of(_219_v39);
        (_195_globalState).f7 = _dafny.Seq.contains((_module.D11.create_DC27(_219_v39, new BigNumber((_230_v50).length), _219_v39, _231_v51, _232_v52)).dtor_cf43, _221_v41);
      }
      _176_v2 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_176_v2));
      let _233_v53;
      let _nw34 = new _module.C1();
      _nw34.__ctor(_174_v0, false, _174_v0, _176_v2);
      _233_v53 = _nw34;
      _233_v53 = _233_v53;
      let _234_v54;
      _234_v54 = _module.D10.create_DC23(_183_v9);
      let _source7 = (((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))]) ? (_234_v54) : (_234_v54));
      if (_source7.is_DC24) {
        let _235___mcc_h0 = (_source7).cf32;
        let _236___mcc_h1 = (_source7).cf33;
        let _237___mcc_h2 = (_source7).cf34;
        let _238___mcc_h3 = (_source7).cf35;
        let _239___mcc_h4 = (_source7).cf36;
        let _240_cf36 = _239___mcc_h4;
        let _241_cf35 = _238___mcc_h3;
        let _242_cf34 = _237___mcc_h2;
        let _243_cf33 = _236___mcc_h1;
        let _244_cf32 = _235___mcc_h0;
        let _245_v55;
        _245_v55 = _module.D3.create_DC6(_186_v12);
        let _246_v56;
        _246_v56 = _dafny.Map.Empty.slice().updateUnsafe((_245_v55).dtor_cf4,_dafny.Seq.Concat(_192_v18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(968)), ((_247_v9) => function (_248_i2) {
          return _247_v9;
        })(_183_v9))));
        _246_v56 = (_246_v56).update(_186_v12, _dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), ((_249_v9) => function (_250_i3) {
          return _249_v9;
        })(_183_v9)));
        let _251_v57;
        _251_v57 = _dafny.Map.Empty.slice().updateUnsafe((_233_v53).f23,_243_cf33);
        let _252_v58;
        let _nw35 = new _module.C0();
        _nw35.__ctor((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))], ((_233_v53).f23) === (_233_v53.f22), _module.__default.safeDivisionInt(_243_cf33, new BigNumber((_251_v57).length)));
        _252_v58 = _nw35;
        let _253_v59;
        let _nw36 = new _module.C1();
        _nw36.__ctor(_233_v53.f22, _240_cf36, _233_v53.f22, _176_v2);
        _253_v59 = _nw36;
        let _254_v60;
        _254_v60 = _dafny.Map.Empty.slice().updateUnsafe((_252_v58).f24,_253_v59);
        let _255_v61;
        _255_v61 = _module.D5.create_DC11((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))], _244_cf32, _233_v53.f22, _192_v18, new BigNumber((_254_v60).length));
        let _nw37 = new _module.C0();
        _nw37.__ctor((_module.__default.fm18(_255_v61, _module.__default.fm12(_240_cf36, false, _233_v53.f22, (_253_v59).f19, _195_globalState), (_253_v59).f19, _195_globalState)).dtor_cf12, !((((_178_v4).contains(new BigNumber((_192_v18).length))) ? ((_178_v4).get(new BigNumber((_192_v18).length))) : (_233_v53.f22))) || ((_252_v58).f24), _176_v2);
        _252_v58 = _nw37;
        let _256_v62;
        let _nw38 = Array((new BigNumber(2)).toNumber());
        _256_v62 = _nw38;
        let _257_v63;
        _257_v63 = _dafny.Seq.of(_256_v62);
        let _258_v64;
        _258_v64 = _dafny.Seq.of(new BigNumber((_192_v18).length), (_253_v59).f19);
        _257_v63 = _dafny.Seq.update(_257_v63, _module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_258_v64).length), _244_cf32), new BigNumber((_257_v63).length)), _256_v62);
        _192_v18 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("igvjgd"), _192_v18);
      } else if (_source7.is_DC23) {
        let _259___mcc_h5 = (_source7).cf31;
        let _260_cf31 = _259___mcc_h5;
        let _index27 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_180_v6).length));
        (_180_v6)[_index27] = _module.__default.safeModuloInt(_176_v2, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-124)), ((_261_v9) => function (_262_i4) {
          return _261_v9;
        })(_183_v9))).length)));
        let _index28 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_180_v6).length));
        (_180_v6)[_index28] = (_176_v2).multipliedBy(new BigNumber(725));
        let _263_v65;
        let _nw39 = new _module.C2();
        _nw39.__ctor(!(_233_v53.f22), _176_v2);
        _263_v65 = _nw39;
        let _264_v66;
        let _nw40 = new _module.C3();
        _nw40.__ctor(_186_v12, _192_v18, !((_176_v2).isLessThanOrEqualTo((_180_v6)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_180_v6).length))])), _176_v2);
        _264_v66 = _nw40;
        let _265_v67;
        let _nw41 = new _module.C3();
        _nw41.__ctor(_186_v12, _192_v18, _174_v0, (_180_v6)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_180_v6).length))]);
        _265_v67 = _nw41;
        let _266_v68;
        _266_v68 = _dafny.Set.fromElements(_265_v67, _265_v67);
        let _267_v69;
        let _init2 = ((_268_v67) => function (_269_i5) {
          return (_269_i5).multipliedBy((_268_v67).f19);
        })(_265_v67);
        let _nw42 = Array((new BigNumber(28)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw42.length); _i0_2++) {
          _nw42[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _267_v69 = _nw42;
        let _270_v70;
        _270_v70 = _dafny.Set.fromElements(_267_v69);
        let _271_v71;
        _271_v71 = _module.D8.create_DC19((_264_v66).f20, (_233_v53).f23, _266_v68, _270_v70, _233_v53.f22);
        let _272_v72;
        _272_v72 = _dafny.Seq.of((_180_v6)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_180_v6).length))]);
        let _273_v73;
        let _nw43 = new _module.C1();
        _nw43.__ctor((_271_v71).dtor_cf25, false, _dafny.Seq.contains(_272_v72, _176_v2), _module.__default.safeModuloInt(_176_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),(_265_v67).f18)).length)));
        _273_v73 = _nw43;
      } else {
        let _274___mcc_h6 = (_source7).cf37;
        let _275_cf37 = _274___mcc_h6;
        _192_v18 = _192_v18;
        let _index29 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
        (_186_v12)[_index29] = !((_233_v53).f23);
        let _index30 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
        (_186_v12)[_index30] = _dafny.Seq.IsProperPrefixOf(_193_v19, _193_v19);
        let _276_v74;
        let _init3 = ((_277_v18) => function (_278_i6) {
          return (_277_v18)[_module.__default.safeIndex(new BigNumber((_277_v18).length), new BigNumber((_277_v18).length))];
        })(_192_v18);
        let _nw44 = Array((new BigNumber(19)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw44.length); _i0_3++) {
          _nw44[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _276_v74 = _nw44;
        let _index31 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_276_v74).length));
        (_276_v74)[_index31] = _183_v9;
        let _index32 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
        let _index33 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_276_v74).length));
        let _index34 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
        let _rhs27 = _174_v0;
        let _rhs28 = (_192_v18)[_module.__default.safeIndex(_176_v2, new BigNumber((_192_v18).length))];
        let _rhs29 = new _dafny.CodePoint('o'.codePointAt(0));
        let _rhs30 = (_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))];
        let _lhs21 = _186_v12;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
        let _lhs23 = _276_v74;
        let _lhs24 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_276_v74).length));
        let _lhs25 = _186_v12;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
        _lhs21[_lhs22] = _rhs27;
        _183_v9 = _rhs28;
        _lhs23[_lhs24] = _rhs29;
        _lhs25[_lhs26] = _rhs30;
      }
      let _279_v75;
      let _nw45 = new _module.C0();
      _nw45.__ctor(!(new BigNumber(-235)).isEqualTo(_176_v2), true, new BigNumber(712));
      _279_v75 = _nw45;
      let _280_v76;
      _280_v76 = _dafny.Seq.of(_180_v6);
      _180_v6 = (_280_v76)[_module.__default.safeIndex(_module.__default.safeModuloInt(_176_v2, _176_v2), new BigNumber((_280_v76).length))];
      (_233_v53).f22 = _dafny.areEqual(_192_v18, _module.__default.fm9(_176_v2, _195_globalState));
      let _281_v77;
      _281_v77 = _dafny.MultiSet.fromElements(_176_v2);
      _186_v12 = (((_194_v20).contains(_module.__default.safeModuloInt(new BigNumber((_module.__default.fm28(_195_globalState)).length), (((_281_v77).contains(new BigNumber(836))) ? ((_281_v77).get(new BigNumber(836))) : (_176_v2))))) ? ((_194_v20).get(_module.__default.safeModuloInt(new BigNumber((_module.__default.fm28(_195_globalState)).length), (((_281_v77).contains(new BigNumber(836))) ? ((_281_v77).get(new BigNumber(836))) : (_176_v2))))) : ((_187_v13)[_module.__default.safeIndex(_176_v2, new BigNumber((_187_v13).length))]));
      let _282_v78;
      _282_v78 = _dafny.Seq.of(new BigNumber((_192_v18).length));
      let _283_v79;
      _283_v79 = _dafny.Map.Empty.slice().updateUnsafe(_183_v9,_282_v78);
      _283_v79 = (_283_v79).Merge(_283_v79);
      _192_v18 = _dafny.Seq.Concat(((_233_v53.f22) ? (_192_v18) : (_192_v18)), _dafny.Seq.Concat(_192_v18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(881)), ((_284_v9) => function (_285_i7) {
        return _284_v9;
      })(_183_v9))));
      let _286_v80;
      _286_v80 = _dafny.Map.Empty.slice().updateUnsafe(_183_v9,_176_v2);
      let _287_v81;
      _287_v81 = _dafny.Map.Empty.slice().updateUnsafe(_286_v80,_176_v2);
      let _288_v82;
      _288_v82 = _dafny.Map.Empty.slice().updateUnsafe(_180_v6,(new BigNumber((_192_v18).length)).isEqualTo(new BigNumber((_287_v81).length)));
      (_195_globalState).f7 = !((((_288_v82).contains(_180_v6)) ? ((_288_v82).get(_180_v6)) : (_174_v0)));
      let _hi1 = new BigNumber((_dafny.Seq.of((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))], (_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))])).length);
      for (let _289_i8 = new BigNumber(542); _289_i8.isLessThan(_hi1); _289_i8 = _289_i8.plus(_dafny.ONE)) {
        if ((_279_v75).f24) {
          (_233_v53).f22 = (_module.__default.safeModuloInt((_282_v78)[_module.__default.safeIndex(_176_v2, new BigNumber((_282_v78).length))], _176_v2)).isEqualTo(_176_v2);
          _192_v18 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("bgw"), _module.__default.safeIndex(_176_v2, new BigNumber((_dafny.Seq.UnicodeFromString("bgw")).length)), _183_v9);
          let _rhs31 = (((_279_v75).f24) ? ((_233_v53).f23) : ((_279_v75).f24));
          let _rhs32 = _233_v53.f22;
          let _rhs33 = (_dafny.ZERO).minus(_module.__default.fm0(_195_globalState));
          let _lhs27 = _233_v53;
          let _lhs28 = _233_v53;
          _lhs27.f22 = _rhs31;
          _lhs28.f22 = _rhs32;
          _176_v2 = _rhs33;
          (_195_globalState).f7 = (_233_v53).f23;
          let _290_v83;
          let _nw46 = new _module.C1();
          _nw46.__ctor((_279_v75).f24, _233_v53.f22, (_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))], _module.__default.safeModuloInt(new BigNumber((_175_v1).length), _289_i8));
          _290_v83 = _nw46;
          let _nw47 = new _module.C1();
          _nw47.__ctor((_233_v53).f23, (_279_v75).f24, _233_v53.f22, new BigNumber(-707));
          _290_v83 = _nw47;
        } else {
          let _291_v84;
          let _nw48 = new _module.C2();
          _nw48.__ctor(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_183_v9, _183_v9), _192_v18), new BigNumber(775));
          _291_v84 = _nw48;
          let _292_v85;
          _292_v85 = _dafny.Seq.of(_291_v84);
          _291_v84 = (_292_v85)[_module.__default.safeIndex(_289_i8, new BigNumber((_292_v85).length))];
          let _293_v86;
          let _nw49 = Array((new BigNumber(29)).toNumber()).fill([]);
          _293_v86 = _nw49;
          let _rhs34 = _293_v86;
          let _rhs35 = (_175_v1);
          _293_v86 = _rhs34;
          _175_v1 = _rhs35;
          _176_v2 = _module.__default.safeModuloInt((_module.__default.fm0(_195_globalState)).plus(_176_v2), new BigNumber(-886));
          let _294_v87;
          _294_v87 = _dafny.MultiSet.fromElements((_233_v53).f23);
          _294_v87 = (_294_v87).Intersect((_294_v87).Union(_294_v87));
          (_195_globalState).f7 = !(!(_233_v53.f22) || (_233_v53.f22));
        }
        let _295_v88;
        _295_v88 = _module.D7.create_DC15((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))]);
        let _296_v89;
        _296_v89 = _dafny.MultiSet.fromElements(_295_v88, _295_v88, _295_v88, _295_v88, _module.D7.create_DC15(_233_v53.f22));
        if ((_296_v89).IsProperSubsetOf(_296_v89)) {
          let _index35 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_180_v6).length));
          (_180_v6)[_index35] = _module.__default.safeModuloInt(_289_i8, new BigNumber((_282_v78).length));
          let _297_v90;
          let _nw50 = Array((new BigNumber(19)).toNumber()).fill([]);
          _297_v90 = _nw50;
          let _298_v91;
          let _nw51 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
          _298_v91 = _nw51;
          let _index36 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_297_v90).length));
          (_297_v90)[_index36] = _298_v91;
          let _index37 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_180_v6).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_297_v90).length));
          let _rhs36 = _176_v2;
          let _rhs37 = _298_v91;
          let _lhs29 = _180_v6;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_180_v6).length));
          let _lhs31 = _297_v90;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_297_v90).length));
          _lhs29[_lhs30] = _rhs36;
          _lhs31[_lhs32] = _rhs37;
          let _index39 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_180_v6).length));
          (_180_v6)[_index39] = _module.__default.safeDivisionInt(_289_i8, new BigNumber((_dafny.Seq.Concat(_282_v78, _282_v78)).length));
          (_233_v53).f22 = _233_v53.f22;
          _192_v18 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qbhnhm"), _192_v18), _192_v18);
          let _299_v92;
          _299_v92 = _dafny.Map.Empty.slice().updateUnsafe(_183_v9,(_279_v75).f24);
          let _300_v93;
          let _nw52 = Array((new BigNumber(2)).toNumber());
          _nw52[(_dafny.ZERO).toNumber()] = _299_v92;
          _nw52[(_dafny.ONE).toNumber()] = _299_v92;
          _300_v93 = _nw52;
          let _301_v95;
          _301_v95 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm16((_279_v75).f24, (_180_v6)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_180_v6).length))], (_279_v75).f24, _195_globalState),new BigNumber((function () {
            let _coll13 = new _dafny.Set();
            for (const _compr_13 of _dafny.IntegerRange(new BigNumber(-52), new BigNumber(102))) {
              let _302_v94 = _compr_13;
              if (((new BigNumber(-52)).isLessThanOrEqualTo(_302_v94)) && ((_302_v94).isLessThan(new BigNumber(102)))) {
                _coll13.add(_module.__default.safeModuloInt(_302_v94, (_180_v6)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_180_v6).length))]));
              }
            }
            return _coll13;
          }()).length));
          let _index40 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_300_v93).length));
          (_300_v93)[_index40] = (_module.__default.fm23(_301_v95, _195_globalState)).Merge(function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of (_192_v18).Elements) {
              let _303_v96 = _compr_14;
              if (_dafny.Seq.contains(_192_v18, _303_v96)) {
                _coll14.push([_303_v96,_233_v53.f22]);
              }
            }
            return _coll14;
          }());
          let _index41 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_300_v93).length));
          (_300_v93)[_index41] = _299_v92;
        } else {
          let _304_v97;
          _304_v97 = _dafny.Seq.of(_233_v53);
          let _index42 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
          let _rhs38 = !_dafny.areEqual(_dafny.Seq.Concat(_304_v97, _304_v97), _304_v97);
          let _rhs39 = (_233_v53).f23;
          let _rhs40 = _233_v53.f22;
          let _lhs33 = _195_globalState;
          let _lhs34 = _186_v12;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
          _lhs33.f7 = _rhs38;
          _174_v0 = _rhs39;
          _lhs34[_lhs35] = _rhs40;
          let _305_v98;
          let _nw53 = new _module.C2();
          _nw53.__ctor(_233_v53.f22, new BigNumber((_dafny.Set.fromElements(_176_v2)).length));
          _305_v98 = _nw53;
          let _306_v99;
          let _307_v100;
          let _308_v101;
          let _out7;
          let _out8;
          let _out9;
          let _outcollector2 = (_233_v53).m5(_305_v98, _195_globalState);
          _out7 = _outcollector2[0];
          _out8 = _outcollector2[1];
          _out9 = _outcollector2[2];
          _306_v99 = _out7;
          _307_v100 = _out8;
          _308_v101 = _out9;
          let _309_v102;
          _309_v102 = _module.D5.create_DC11(_module.__default.fm16(_233_v53.f22, _307_v100, _233_v53.f22, _195_globalState), _307_v100, (_279_v75).f24, _192_v18, new BigNumber((_dafny.Seq.update(_192_v18, _module.__default.safeIndex(new BigNumber((_286_v80).length), new BigNumber((_192_v18).length)), _183_v9)).length));
          let _310_v103;
          let _nw54 = new _module.C3();
          _nw54.__ctor(_186_v12, _192_v18, _233_v53.f22, _176_v2);
          _310_v103 = _nw54;
          let _311_v104;
          _311_v104 = _dafny.Seq.of(_310_v103);
          let _312_v105;
          _312_v105 = _dafny.Set.fromElements(_233_v53.f22, !((_186_v12)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length))]));
          _183_v9 = _module.__default.fm12((_309_v102).dtor_cf14, (_279_v75).f24, (_233_v53).f23, (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_311_v104).length), new BigNumber((_312_v105).length))), _195_globalState);
          (_233_v53).f22 = (_279_v75).f24;
          let _313_v106;
          _313_v106 = _dafny.Map.Empty.slice().updateUnsafe(_233_v53.f22,_308_v101);
          let _314_v107;
          _314_v107 = _dafny.Seq.of((_313_v106).Merge((_313_v106).update(_233_v53.f22, _308_v101)), (_313_v106).Merge(_dafny.Map.Empty.slice().updateUnsafe((_233_v53).f23,(_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_310_v103).f21, (_310_v103).f21)).length)))), _313_v106);
          _314_v107 = _dafny.Seq.of(_313_v106, _313_v106, _313_v106, _313_v106, (_313_v106).Merge(_313_v106));
        }
        let _315_v108;
        _315_v108 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Concat(_175_v1, _175_v1));
        _176_v2 = new BigNumber((_315_v108).length);
        let _index43 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_186_v12).length));
        (_186_v12)[_index43] = ((_279_v75).f24) === (_233_v53.f22);
      }
      process.stdout.write(_dafny.toString(_174_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_175_v1, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_176_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[_dafny.ZERO], _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[_dafny.ONE], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(2)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(3)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(4)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(5)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(6)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(7)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(8)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(9)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(10)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(11)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(12)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(13)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(14)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(15)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(16)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(17)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(18)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(19)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(20)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(21)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(22)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(23)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(24)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(25)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(26)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_177_v3)[new BigNumber(27)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_179_v5), _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v6)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_v7).dtor_cf1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(3)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(4)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(5)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(6)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(7)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(8)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(9)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(10)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(11)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(12)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(13)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(14)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(15)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(16)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(17)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(18)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(19)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(20)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(21)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(22)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(23)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(24)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(25)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(26)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v8)[new BigNumber(27)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_183_v9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v12)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_187_v13).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v15).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))).updateUnsafe(new BigNumber(2),new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,new _dafny.CodePoint('o'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_191_v17).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_192_v18).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_193_v19, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ldynryh"), _dafny.Seq.UnicodeFromString("hf")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_194_v20).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[_dafny.ZERO], _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[_dafny.ONE], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(2)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(3)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(4)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(5)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(6)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(7)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(8)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(9)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(10)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(11)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(12)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(13)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(14)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(15)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(16)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(17)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(18)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(19)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(20)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(21)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(22)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(23)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(24)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(25)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(26)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_195_globalState).f0)[new BigNumber(27)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(3)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(4)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(5)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(6)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(7)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(8)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(9)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(10)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(11)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(12)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(13)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(14)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(15)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(16)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(17)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(18)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(19)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(20)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(21)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(22)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(23)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(24)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(25)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(26)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_195_globalState).f1)[new BigNumber(27)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_195_globalState).f3).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new _dafny.CodePoint('o'.codePointAt(0))).updateUnsafe(new BigNumber(2),new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,new _dafny.CodePoint('o'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState.f6).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_195_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_195_globalState).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_195_globalState).f15, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ldynryh"), _dafny.Seq.UnicodeFromString("hf")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_195_globalState).f16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_200_v24).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_233_v53.f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_233_v53).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v54).dtor_cf31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_279_v75).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_280_v76).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v77).equals(_dafny.MultiSet.fromElements(new BigNumber(173)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_282_v78, _dafny.Seq.of(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_283_v79).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),_dafny.Seq.of(new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_286_v80).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber(173)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_287_v81).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber(173)),new BigNumber(173)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_288_v82).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2() {
      let $dt = new D1(0);
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf1 === other.cf1;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4() {
      let $dt = new D2(0);
      return $dt;
    }
    static create_DC3(cf2) {
      let $dt = new D2(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC5(cf3) {
      let $dt = new D2(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4";
      } else if (this.$tag === 1) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf5, cf6, cf7, cf8, cf9) {
      let $dt = new D3(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC6(cf4) {
      let $dt = new D3(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf4 === other.cf4;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7(null, _dafny.Map.Empty, false, _dafny.Set.Empty, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC8(cf10) {
      let $dt = new D4(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9";
      } else if (this.$tag === 1) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC9();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf12, cf13, cf14, cf15, cf16) {
      let $dt = new D5(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC10(cf11) {
      let $dt = new D5(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + this.cf15.toVerbatimString(true) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12 && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf11 === other.cf11;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC11(false, _dafny.ZERO, false, _dafny.Seq.UnicodeFromString(""), _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC13(cf18) {
      let $dt = new D6(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC12(cf17) {
      let $dt = new D6(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC13(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf20) {
      let $dt = new D7(0);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC14(cf19) {
      let $dt = new D7(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC16(cf21) {
      let $dt = new D7(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC15(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC18(cf23) {
      let $dt = new D8(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC19(cf24, cf25, cf26, cf27, cf28) {
      let $dt = new D8(1);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC17(cf22) {
      let $dt = new D8(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24 && this.cf25 === other.cf25 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC18(_dafny.Map.Empty);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21() {
      let $dt = new D9(0);
      return $dt;
    }
    static create_DC20(cf29) {
      let $dt = new D9(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC22(cf30) {
      let $dt = new D9(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC21";
      } else if (this.$tag === 1) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC21();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC24(cf32, cf33, cf34, cf35, cf36) {
      let $dt = new D10(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC23(cf31) {
      let $dt = new D10(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC25(cf37) {
      let $dt = new D10(2);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC24" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC25" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35 && this.cf36 === other.cf36;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC24(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC27(cf39, cf40, cf41, cf42, cf43) {
      let $dt = new D11(0);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC26(cf38) {
      let $dt = new D11(1);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf39, other.cf39) && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42) && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC27(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.MultiSet.Empty, _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28(cf44) {
      let $dt = new D12(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf46, cf47) {
      let $dt = new D13(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC31(cf48, cf49, cf50) {
      let $dt = new D13(1);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC29(cf45) {
      let $dt = new D13(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC32(cf51) {
      let $dt = new D13(3);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get is_DC32() { return this.$tag === 3; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC31" + "(" + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC32" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf48 === other.cf48 && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC30(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f6 = _dafny.Map.Empty;
      this.f7 = false;
      this._f0 = [];
      this._f1 = [];
      this._f2 = _dafny.ZERO;
      this._f3 = _dafny.MultiSet.Empty;
      this._f4 = _dafny.ZERO;
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.Map.Empty;
      this._f10 = false;
      this._f11 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
      this._f13 = false;
      this._f14 = false;
      this._f15 = _dafny.Seq.of();
      this._f16 = _dafny.Map.Empty;
      this._f17 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f18 = false;
      this._f19 = _dafny.ZERO;
      this._f24 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    __ctor(f24, f18, f19) {
      let _this = this;
      (_this)._f24 = f24;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm8(p0, p1, globalState) {
      let _this = this;
      return _module.D1.create_DC2();
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f18 = false;
      this._f19 = _dafny.ZERO;
      this.f22 = false;
      this._f23 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    __ctor(f22, f23, f18, f19) {
      let _this = this;
      (_this).f22 = f22;
      (_this)._f23 = f23;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      if (false) {
        return _module.D2.create_DC3(_dafny.Seq.of(_module.D1.create_DC2(), _module.D1.create_DC2(), _module.D1.create_DC2()));
      } else {
        return _module.D2.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(147)), function (_316_i0) {
  return _module.D1.create_DC2();
}));
      }
    };
    fm7(p0, p1, globalState) {
      let _this = this;
      return !(!((_this).f18) || (!((_this).f18))) || ((_this).f18);
    };
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.Map.Empty;
      let _317_v0;
      _317_v0 = new BigNumber(117);
      let _318_v1;
      _318_v1 = p0;
      let _pat_let_tv12 = _317_v0;
      _317_v0 = function (_source8) {
        let _319___mcc_h0 = _source8;
        let _320_cf0 = _319___mcc_h0;
        return _pat_let_tv12;
      }(_318_v1);
      let _321_v2;
      let _nw55 = new _module.C0();
      _nw55.__ctor(p1, false, (_this).f19);
      _321_v2 = _nw55;
      let _322_v3;
      _322_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,(_this).f19);
      let _323_v4;
      _323_v4 = _module.D3.create_DC7(_321_v2, _322_v3, true, _dafny.Set.fromElements(p1), (_this).f23);
      let _324_v5;
      _324_v5 = _dafny.Set.fromElements((_this).f18, _this.f22, (_321_v2).f18);
      let _325_v6;
      let _nw56 = Array((new BigNumber(9)).toNumber());
      _nw56[(_dafny.ZERO).toNumber()] = _this.f22;
      _nw56[(_dafny.ONE).toNumber()] = true;
      _nw56[(new BigNumber(2)).toNumber()] = (_323_v4).dtor_cf9;
      _nw56[(new BigNumber(3)).toNumber()] = (_this).f18;
      _nw56[(new BigNumber(4)).toNumber()] = (_324_v5).IsProperSubsetOf(_dafny.Set.fromElements(false));
      _nw56[(new BigNumber(5)).toNumber()] = (_dafny.Set.fromElements((_this).f18, (_this).f18)).IsProperSubsetOf(_dafny.Set.fromElements(false, _this.f22, p1));
      _nw56[(new BigNumber(6)).toNumber()] = ((false) ? ((_321_v2).f18) : ((_321_v2).f18));
      _nw56[(new BigNumber(7)).toNumber()] = (_this).f18;
      _nw56[(new BigNumber(8)).toNumber()] = true;
      _325_v6 = _nw56;
      let _index44 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_325_v6).length));
      (_325_v6)[_index44] = (_this).f23;
      let _index45 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_325_v6).length));
      (_325_v6)[_index45] = !(!(!(((_dafny.ZERO).minus(_317_v0)).isLessThan((_321_v2).f19))));
      let _326_v7;
      let _init4 = ((_327_p0, _328_p1) => function (_329_i0) {
        return _module.__default.safeModuloInt(_329_i0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_327_p0,_328_p1)).length));
      })(p0, p1);
      let _nw57 = Array((new BigNumber(29)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw57.length); _i0_4++) {
        _nw57[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _326_v7 = _nw57;
      let _index46 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_326_v7).length));
      (_326_v7)[_index46] = _module.__default.fm0(globalState);
      let _330_v8;
      _330_v8 = _dafny.MultiSet.fromElements((_321_v2).f18);
      let _index47 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_325_v6).length));
      let _index48 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_326_v7).length));
      let _rhs41 = ((_this).f18) && (false);
      let _rhs42 = _module.__default.safeDivisionInt(_317_v0, ((_this).f19).plus(new BigNumber(37)));
      let _rhs43 = (((_dafny.MultiSet.fromElements((_321_v2).f18, (_this).f23)).IsProperSubsetOf(_330_v8)) ? (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(238)), function (_331_i1) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length)) : ((_this).f19));
      let _lhs36 = _325_v6;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_325_v6).length));
      let _lhs38 = _326_v7;
      let _lhs39 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_326_v7).length));
      _lhs36[_lhs37] = _rhs41;
      _317_v0 = _rhs42;
      _lhs38[_lhs39] = _rhs43;
      let _332_v9;
      _332_v9 = _dafny.Map.Empty.slice().updateUnsafe(_321_v2,p1);
      let _333_v10;
      _333_v10 = _dafny.Seq.of(new BigNumber((_332_v9).length));
      let _334_v11;
      _334_v11 = _dafny.Seq.UnicodeFromString("nr");
      let _335_i2;
      _335_i2 = _dafny.ZERO;
      L2: {
        while (!(_module.__default.safeModuloInt((_333_v10)[_module.__default.safeIndex((_326_v7)[_module.__default.safeIndex(new BigNumber(868), new BigNumber((_326_v7).length))], new BigNumber((_333_v10).length))], (_this).f19)).isEqualTo(new BigNumber((_dafny.Seq.Concat(_334_v11, _334_v11)).length))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_335_i2)) {
              break L2;
            }
            _335_i2 = (_335_i2).plus(_dafny.ONE);
            let _index49 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_325_v6).length));
            (_325_v6)[_index49] = !(((_dafny.Seq.IsPrefixOf(_334_v11, _334_v11)) ? (!_dafny.areEqual(_334_v11, _334_v11)) : (_dafny.Seq.IsPrefixOf(_334_v11, _dafny.Seq.Create(_module.__default.abs(new BigNumber(411)), function (_336_i3) {
              return new _dafny.CodePoint('e'.codePointAt(0));
            })))));
            let _index50 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_326_v7).length));
            (_326_v7)[_index50] = (_dafny.ZERO).minus(((_317_v0).multipliedBy((_this).f19)).plus((_this).f19));
            let _337_v12;
            _337_v12 = new _dafny.CodePoint('x'.codePointAt(0));
            let _338_v13;
            let _nw58 = new _module.C0();
            _nw58.__ctor((_325_v6)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_325_v6).length))], _dafny.Seq.contains(_module.__default.fm9((_321_v2).f19, globalState), _337_v12), new BigNumber(581));
            _338_v13 = _nw58;
            _338_v13 = _338_v13;
            let _339_v14;
            let _nw59 = new _module.C0();
            _nw59.__ctor((_this.f22) === (p1), (_this).f18, (_this).f19);
            _339_v14 = _nw59;
          }
        }
      }
      let _index51 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_326_v7).length));
      (_326_v7)[_index51] = ((_326_v7)[_module.__default.safeIndex(new BigNumber(868), new BigNumber((_326_v7).length))]).plus(((_326_v7)[_module.__default.safeIndex(new BigNumber(868), new BigNumber((_326_v7).length))]).plus((_this).f19));
      (globalState).f7 = p1;
      let _340_v15;
      _340_v15 = new _dafny.CodePoint('e'.codePointAt(0));
      r0 = _340_v15;
      let _341_v16;
      _341_v16 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f23) ? ((_326_v7)[_module.__default.safeIndex(new BigNumber(868), new BigNumber((_326_v7).length))]) : (new BigNumber(553))),_module.D2.create_DC5(_module.D2.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), function (_342_i4) {
  return _module.D1.create_DC2();
}))));
      r1 = _341_v16;
      return [r0, r1];
    }
    m5(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _343_v0;
      _343_v0 = _module.D2.create_DC4();
      let _344_v1;
      _344_v1 = new _dafny.CodePoint('g'.codePointAt(0));
      let _345_v2;
      _345_v2 = _dafny.Set.fromElements((_this).fm7(_343_v0, _344_v1, globalState), (p0).f18, ((_this).f18) || ((_this).f18), (_this).f18, _this.f22);
      let _346_v3;
      _346_v3 = _dafny.Set.fromElements((_this).f19, (_this).f19, (_this).f19, (_this).f19);
      let _347_v4;
      _347_v4 = _dafny.Seq.of(_346_v3);
      _345_v2 = (_module.__default.fm10((p0).f19, new BigNumber((_347_v4).length), globalState)).Intersect((_dafny.Set.fromElements(false)).Union(_345_v2));
      if (((p0).f19).isLessThanOrEqualTo((_this).f19)) {
        let _348_v5;
        _348_v5 = _dafny.Seq.UnicodeFromString("bccgooiv");
        if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(260)), ((_349_v1) => function (_350_i0) {
          return _349_v1;
        })(_344_v1)), _348_v5), _344_v1)) {
          let _351_v6;
          _351_v6 = _dafny.Map.Empty.slice().updateUnsafe((p0).f18,(((_this).f23) ? ((p0).f19) : ((_this).f19)));
          _351_v6 = (_351_v6).update(false, (p0).f19);
          let _352_v7;
          _352_v7 = _dafny.Seq.of((p0).f18, (_this).f18, (p0).f18);
          (globalState).f7 = (_352_v7)[_module.__default.safeIndex((p0).f19, new BigNumber((_352_v7).length))];
          let _353_v8;
          _353_v8 = _dafny.Map.Empty.slice().updateUnsafe((p0).f18,(_this).fm7(_343_v0, _344_v1, globalState));
          _353_v8 = (_353_v8).update((p0).f18, (p0).f18);
          let _354_v9;
          _354_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,new BigNumber((_348_v5).length));
          let _355_v10;
          _355_v10 = _dafny.MultiSet.fromElements((p0).f18);
          let _rhs44 = _module.__default.safeModuloInt((new BigNumber((_348_v5).length)).minus((((_354_v9).contains((p0).f19)) ? ((_354_v9).get((p0).f19)) : ((_dafny.ZERO).minus((p0).f19)))), (p0).f19);
          let _rhs45 = new BigNumber((((_355_v10).Union(_355_v10)).Union(_355_v10)).cardinality());
          r1 = _rhs44;
          r2 = _rhs45;
          let _356_v11;
          _356_v11 = _dafny.MultiSet.fromElements((p0).f19, (p0).f19);
          let _357_v12;
          let _nw60 = Array((new BigNumber(9)).toNumber());
          _nw60[(_dafny.ZERO).toNumber()] = (_this).f19;
          _nw60[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((p0).f18,new BigNumber(((_module.__default.fm11(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_348_v5).length),(p0).f19)).length), true, _module.__default.fm12((p0).f18, (p0).f18, (_this).f23, new BigNumber(-835), globalState), globalState)).update(new BigNumber((_346_v3).length), _module.__default.abs((p0).f19))).cardinality()))).length);
          _nw60[(new BigNumber(2)).toNumber()] = (new BigNumber((_351_v6).length)).plus((_this).f19);
          _nw60[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((_this).f19, (_this).f19);
          _nw60[(new BigNumber(4)).toNumber()] = (p0).f19;
          _nw60[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((p0).f19);
          _nw60[(new BigNumber(6)).toNumber()] = (p0).f19;
          _nw60[(new BigNumber(7)).toNumber()] = new BigNumber((_352_v7).length);
          _nw60[(new BigNumber(8)).toNumber()] = new BigNumber((_356_v11).cardinality());
          _357_v12 = _nw60;
          let _nw61 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _357_v12 = _nw61;
        } else {
          let _358_v13;
          let _init5 = ((_359_v5, _360_v1) => function (_361_i1) {
            return _dafny.areEqual(_359_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(240)), ((_362_v1) => function (_363_i2) {
              return _362_v1;
            })(_360_v1)));
          })(_348_v5, _344_v1);
          let _nw62 = Array((new BigNumber(4)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw62.length); _i0_5++) {
            _nw62[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _358_v13 = _nw62;
          let _index52 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_358_v13).length));
          (_358_v13)[_index52] = (_this).f23;
          let _364_v14;
          _364_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_dafny.ZERO).minus((p0).f19));
          let _365_v15;
          _365_v15 = _module.D3.create_DC7(p0, _364_v14, (p0).f18, _345_v2, true);
          let _366_v16;
          _366_v16 = _dafny.Seq.of(_365_v15, _365_v15, _365_v15, _365_v15, _365_v15);
          let _367_v17;
          _367_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm7(_343_v0, _344_v1, globalState),!_dafny.areEqual(_366_v16, _366_v16));
          let _index53 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_358_v13).length));
          (_358_v13)[_index53] = (((_367_v17).contains((_this).f23)) ? ((_367_v17).get((_this).f23)) : ((_this).fm7(_343_v0, _344_v1, globalState)));
          let _368_v18;
          _368_v18 = _dafny.Set.fromElements(_348_v5, _348_v5, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(825)), ((_369_v1) => function (_370_i3) {
            return _369_v1;
          })(_344_v1)), _348_v5));
          _368_v18 = _module.__default.fm13(globalState);
          let _index54 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_358_v13).length));
          (_358_v13)[_index54] = (p0).f18;
          (globalState).f7 = (_this).f18;
          let _371_v19;
          _371_v19 = _dafny.Map.Empty.slice().updateUnsafe((p0).f19,_367_v17);
          (globalState).f7 = ((!(!(!(_371_v19).contains(_module.__default.fm0(globalState))))) ? ((_this).f23) : (!((p0).f18)));
        }
        let _372_v20;
        _372_v20 = _module.D3.create_DC7(p0, _dafny.Map.Empty.slice().updateUnsafe(_this.f22,(_this).f19), !((_this).f23), _345_v2, (p0).f18);
        let _rhs46 = ((p0).f19).minus((_this).f19);
        let _rhs47 = ((_372_v20).dtor_cf9) === (_this.f22);
        let _rhs48 = !(!((_this).f18));
        let _rhs49 = (p0).f18;
        let _lhs40 = globalState;
        let _lhs41 = _this;
        let _lhs42 = _this;
        r1 = _rhs46;
        _lhs40.f7 = _rhs47;
        _lhs41.f22 = _rhs48;
        _lhs42.f22 = _rhs49;
        (globalState).f7 = false;
        let _373_v21;
        _373_v21 = _dafny.Seq.of((_this).f23);
        let _374_v22;
        _374_v22 = _dafny.Seq.of(_373_v21);
        let _375_v24;
        _375_v24 = _dafny.Set.fromElements(_dafny.Seq.update(_373_v21, _module.__default.safeIndex(new BigNumber(342), new BigNumber((_373_v21).length)), (_this).f23), _373_v21, _373_v21);
        let _376_v25;
        let _377_v26;
        let _out10;
        let _out11;
        let _outcollector3 = (_this).m4(_dafny.Seq.Concat((_374_v22)[_module.__default.safeIndex((p0).f19, new BigNumber((_374_v22).length))], _373_v21), (((p0).f18) ? (!((_this).f23)) : ((_this).f23)), function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of (_375_v24).Elements) {
            let _378_v23 = _compr_15;
            if ((_375_v24).contains(_378_v23)) {
              _coll15.push([_378_v23,_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_379_v1) => function (_380_i4) {
                return _379_v1;
              })(_344_v1))]);
            }
          }
          return _coll15;
        }(), globalState);
        _out10 = _outcollector3[0];
        _out11 = _outcollector3[1];
        _376_v25 = _out10;
        _377_v26 = _out11;
        let _381_v27;
        _381_v27 = _dafny.Seq.of((_this).f19);
        let _382_v28;
        _382_v28 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_381_v27).length),_this.f22));
        let _383_v29;
        _383_v29 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18) || ((p0).f18),(_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_382_v28).length), (p0).f19)));
        _383_v29 = (_383_v29).update(((_this).f19).isLessThan((p0).f19), (_this).f19);
      } else {
        let _384_v30;
        _384_v30 = _dafny.Seq.of(_module.__default.fm0(globalState), (p0).f19, (p0).f19);
        let _385_v31;
        _385_v31 = _dafny.MultiSet.fromElements((_this).f18);
        let _386_v32;
        _386_v32 = _dafny.Seq.of(_385_v31);
        r2 = (_384_v30)[_module.__default.safeIndex(((_this).f19).minus(new BigNumber((_386_v32).length)), new BigNumber((_384_v30).length))];
        r2 = (_dafny.ZERO).minus((_this).f19);
        let _387_v33;
        let _nw63 = new _module.C0();
        _nw63.__ctor((p0).f18, (_this).f23, new BigNumber((_dafny.Seq.UnicodeFromString("k")).length));
        _387_v33 = _nw63;
        let _388_v34;
        _388_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-762),_387_v33);
        _388_v34 = _388_v34;
        let _389_v35;
        let _nw64 = Array((new BigNumber(2)).toNumber());
        _nw64[(_dafny.ZERO).toNumber()] = (p0).f19;
        _nw64[(_dafny.ONE).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ogkumllt")).length))).minus((p0).f19);
        _389_v35 = _nw64;
        let _index55 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_389_v35).length));
        (_389_v35)[_index55] = (p0).f19;
        let _index56 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_389_v35).length));
        (_389_v35)[_index56] = new BigNumber(-186);
        let _390_v36;
        let _nw65 = Array((new BigNumber(6)).toNumber()).fill([]);
        _390_v36 = _nw65;
        let _391_v37;
        let _nw66 = Array((new BigNumber(26)).toNumber()).fill(false);
        _391_v37 = _nw66;
        let _392_v38;
        _392_v38 = _module.D3.create_DC6(_391_v37);
        let _pat_let_tv13 = _391_v37;
        let _393_v39;
        let _nw67 = Array((new BigNumber(16)).toNumber());
        _nw67[(_dafny.ZERO).toNumber()] = _392_v38;
        _nw67[(_dafny.ONE).toNumber()] = _module.D3.create_DC6(_391_v37);
        _nw67[(new BigNumber(2)).toNumber()] = _392_v38;
        _nw67[(new BigNumber(3)).toNumber()] = _module.D3.create_DC6(_391_v37);
        _nw67[(new BigNumber(4)).toNumber()] = _392_v38;
        _nw67[(new BigNumber(5)).toNumber()] = _392_v38;
        _nw67[(new BigNumber(6)).toNumber()] = _module.D3.create_DC6(_391_v37);
        _nw67[(new BigNumber(7)).toNumber()] = _392_v38;
        _nw67[(new BigNumber(8)).toNumber()] = function (_pat_let2_0) {
          return function (_394_dt__update__tmp_h0) {
            return function (_pat_let3_0) {
              return function (_395_dt__update_hcf4_h0) {
                return _module.D3.create_DC6(_395_dt__update_hcf4_h0);
              }(_pat_let3_0);
            }(_pat_let_tv13);
          }(_pat_let2_0);
        }(_392_v38);
        _nw67[(new BigNumber(9)).toNumber()] = _392_v38;
        _nw67[(new BigNumber(10)).toNumber()] = _module.D3.create_DC6(_391_v37);
        _nw67[(new BigNumber(11)).toNumber()] = _392_v38;
        _nw67[(new BigNumber(12)).toNumber()] = _392_v38;
        _nw67[(new BigNumber(13)).toNumber()] = _392_v38;
        _nw67[(new BigNumber(14)).toNumber()] = _392_v38;
        _nw67[(new BigNumber(15)).toNumber()] = _392_v38;
        _393_v39 = _nw67;
        let _index57 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_390_v36).length));
        (_390_v36)[_index57] = _393_v39;
        let _index58 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_390_v36).length));
        (_390_v36)[_index58] = _393_v39;
      }
      let _396_v40;
      let _nw68 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _396_v40 = _nw68;
      let _index59 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_396_v40).length));
      (_396_v40)[_index59] = (p0).f19;
      let _397_v41;
      _397_v41 = _dafny.Map.Empty.slice().updateUnsafe((p0).f18,(_this).f19);
      let _index60 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_396_v40).length));
      (_396_v40)[_index60] = new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("bomwi"), _module.__default.safeIndex((((_397_v41).contains((_this).f23)) ? ((_397_v41).get((_this).f23)) : (new BigNumber((_dafny.Seq.of(_this.f22, _this.f22)).length))), new BigNumber((_dafny.Seq.UnicodeFromString("bomwi")).length)), _344_v1)).length);
      let _398_v42;
      _398_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(p0).f18);
      let _399_v43;
      _399_v43 = _dafny.Map.Empty.slice().updateUnsafe(_398_v42,(_this).f19);
      let _400_v44;
      _400_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19);
      let _401_v45;
      _401_v45 = _dafny.Map.Empty.slice().updateUnsafe((p0).f19,!(_this.f22));
      let _402_v46;
      let _nw69 = new _module.C0();
      _nw69.__ctor((p0).f18, (_this).fm7(_343_v0, _344_v1, globalState), (((_399_v43).contains(_398_v42)) ? ((_399_v43).get(_398_v42)) : ((((_400_v44).contains((_this).f19)) ? ((_400_v44).get((_this).f19)) : (new BigNumber((_401_v45).length))))));
      _402_v46 = _nw69;
      let _index61 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_396_v40).length));
      let _index62 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_396_v40).length));
      let _rhs50 = _module.__default.fm0(globalState);
      let _rhs51 = (p0).f19;
      let _rhs52 = (_this).f19;
      let _rhs53 = _402_v46;
      let _lhs43 = _396_v40;
      let _lhs44 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_396_v40).length));
      let _lhs45 = _396_v40;
      let _lhs46 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_396_v40).length));
      _lhs43[_lhs44] = _rhs50;
      _lhs45[_lhs46] = _rhs51;
      r1 = _rhs52;
      _402_v46 = _rhs53;
      let _403_i5;
      _403_i5 = _dafny.ZERO;
      L3: {
        while ((_this).f23) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_403_i5)) {
              break L3;
            }
            _403_i5 = (_403_i5).plus(_dafny.ONE);
            r2 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), ((_404_p0) => function (_405_i6) {
              return (_404_p0).f19;
            })(p0))).length);
            let _406_v47;
            let _nw70 = new _module.C0();
            _nw70.__ctor((p0).f18, (_this).f18, (_this).f19);
            _406_v47 = _nw70;
            r2 = _module.__default.safeModuloInt((_396_v40)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_396_v40).length))], (p0).f19);
            (globalState).f7 = ((p0).f18) === ((p0).f18);
          }
        }
      }
      let _407_v48;
      _407_v48 = _dafny.Seq.of((p0).f19);
      let _408_v49;
      _408_v49 = _dafny.Seq.UnicodeFromString("qpxipme");
      let _409_v50;
      _409_v50 = _dafny.Seq.of(true, (_this).f23);
      let _rhs54 = new BigNumber((_407_v48).length);
      let _rhs55 = ((_this.f22) ? (new BigNumber((_dafny.Seq.UnicodeFromString("qoqng")).length)) : ((_this).f19));
      let _rhs56 = _dafny.Seq.IsPrefixOf(_dafny.Seq.of((p0).f18), _dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm14(new BigNumber((_408_v49).length), (p0).f18, globalState), _409_v50), _module.__default.safeIndex((_396_v40)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_396_v40).length))], new BigNumber((_dafny.Seq.Concat(_module.__default.fm14(new BigNumber((_408_v49).length), (p0).f18, globalState), _409_v50)).length)), (_402_v46).f24));
      let _rhs57 = ((_402_v46).f24) || ((_402_v46).f24);
      let _lhs47 = _this;
      let _lhs48 = _this;
      r2 = _rhs54;
      r1 = _rhs55;
      _lhs47.f22 = _rhs56;
      _lhs48.f22 = _rhs57;
      let _410_v51;
      _410_v51 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_396_v40)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_396_v40).length))]));
      let _411_v52;
      _411_v52 = _dafny.Seq.of(_410_v51, _410_v51);
      let _412_v53;
      _412_v53 = _module.D5.create_DC10(_402_v46);
      let _413_v54;
      _413_v54 = _dafny.Seq.of((_412_v53).dtor_cf11);
      let _414_v55;
      _414_v55 = _module.D4.create_DC8(new BigNumber(((_411_v52)[_module.__default.safeIndex((((_410_v51).contains((_this).f19)) ? ((_410_v51).get((_this).f19)) : (new BigNumber((_dafny.Seq.of(new BigNumber((_413_v54).length), new BigNumber(680))).length))), new BigNumber((_411_v52).length))]).cardinality()));
      let _index63 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_396_v40).length));
      (_396_v40)[_index63] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_414_v55).dtor_cf10, (_396_v40)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_396_v40).length))]));
      let _415_v56;
      let _nw71 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Set.Empty);
      _415_v56 = _nw71;
      r0 = _415_v56;
      let _416_v57;
      _416_v57 = _module.D1.create_DC2();
      let _417_v58;
      _417_v58 = _dafny.Seq.of(_416_v57);
      let _418_v59;
      _418_v59 = _module.D2.create_DC5(_module.D2.create_DC3(_417_v58));
      let _419_v60;
      _419_v60 = _module.D2.create_DC5(_418_v59);
      let _pat_let_tv14 = p0;
      let _pat_let_tv15 = p0;
      r1 = function (_source9) {
        if (_source9.is_DC4) {
          return (_this).f19;
        } else if (_source9.is_DC3) {
          let _420___mcc_h0 = (_source9).cf2;
          let _421_cf2 = _420___mcc_h0;
          return (_pat_let_tv14).f19;
        } else {
          let _422___mcc_h1 = (_source9).cf3;
          let _423_cf3 = _422___mcc_h1;
          return (_pat_let_tv15).f19;
        }
      }(_module.D2.create_DC5(_418_v59));
      r2 = (new BigNumber((_410_v51).cardinality())).plus((_407_v48)[_module.__default.safeIndex((p0).f19, new BigNumber((_407_v48).length))]);
      return [r0, r1, r2];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f18 = false;
      this._f19 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    __ctor(f18, f19) {
      let _this = this;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return new _dafny.CodePoint('a'.codePointAt(0));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(!(false), (_this).f18)), _dafny.Seq.Concat(_dafny.Seq.of((_this).f18, (_this).f18), _dafny.Seq.of((_this).f18, (_this).f18))));
    };
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.ZERO;
      if ((_this).f18) {
        r3 = _module.__default.fm0(globalState);
        let _424_v0;
        let _init6 = function (_425_i0) {
          return (_this).f18;
        };
        let _nw72 = Array((new BigNumber(6)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw72.length); _i0_6++) {
          _nw72[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _424_v0 = _nw72;
        let _index64 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_424_v0).length));
        (_424_v0)[_index64] = false;
        let _index65 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_424_v0).length));
        (_424_v0)[_index65] = true;
        r3 = (_this).f19;
        let _426_v1;
        _426_v1 = _dafny.MultiSet.fromElements(p3);
        let _427_v2;
        _427_v2 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(p3, (((_426_v1).contains((_this).f19)) ? ((_426_v1).get((_this).f19)) : ((_this).f19)))).length));
        let _428_v3;
        _428_v3 = _module.D1.create_DC2();
        let _429_v4;
        let _430_v5;
        let _431_v6;
        let _432_v7;
        let _out12;
        let _out13;
        let _out14;
        let _out15;
        let _outcollector4 = (_this).m3(true, new BigNumber((_427_v2).length), p1, _428_v3, globalState);
        _out12 = _outcollector4[0];
        _out13 = _outcollector4[1];
        _out14 = _outcollector4[2];
        _out15 = _outcollector4[3];
        _429_v4 = _out12;
        _430_v5 = _out13;
        _431_v6 = _out14;
        _432_v7 = _out15;
        let _433_v8;
        _433_v8 = _dafny.Set.fromElements(_module.__default.fm0(globalState), new BigNumber(-696));
        let _434_v9;
        let _435_v10;
        let _436_v11;
        let _437_v12;
        let _out16;
        let _out17;
        let _out18;
        let _out19;
        let _outcollector5 = (_this).m3(!(_433_v8).equals(_dafny.Set.fromElements((_this).f19)), p3, (_424_v0)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_424_v0).length))], _428_v3, globalState);
        _out16 = _outcollector5[0];
        _out17 = _outcollector5[1];
        _out18 = _outcollector5[2];
        _out19 = _outcollector5[3];
        _434_v9 = _out16;
        _435_v10 = _out17;
        _436_v11 = _out18;
        _437_v12 = _out19;
      } else {
        let _438_v13;
        let _nw73 = new _module.C0();
        _nw73.__ctor(p2, (_this).f18, _module.__default.fm0(globalState));
        _438_v13 = _nw73;
        let _439_v14;
        let _init7 = function (_440_i1) {
          return _module.__default.safeDivisionInt(_440_i1, (_this).f19);
        };
        let _nw74 = Array((new BigNumber(22)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw74.length); _i0_7++) {
          _nw74[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _439_v14 = _nw74;
        let _index66 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_439_v14).length));
        (_439_v14)[_index66] = p3;
        let _index67 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_439_v14).length));
        (_439_v14)[_index67] = _module.__default.fm0(globalState);
        let _441_v15;
        let _nw75 = new _module.C1();
        _nw75.__ctor(!((_438_v13).f24), (_this).f18, true, p3);
        _441_v15 = _nw75;
        r2 = !(p1) || ((_this).f18);
        (globalState).f7 = p2;
      }
      (globalState).f7 = (_this).f18;
      if ((_this).f18) {
        r3 = ((_dafny.ZERO).minus(((_this).f19).multipliedBy((_this).f19))).plus(new BigNumber(235));
        r2 = p2;
        let _442_v16;
        let _nw76 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
        _442_v16 = _nw76;
        let _443_v17;
        _443_v17 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("heggdukvg")).length)), p3);
        let _index68 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_442_v16).length));
        (_442_v16)[_index68] = _dafny.Seq.Concat(_443_v17, _dafny.Seq.update(_443_v17, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-658)), ((_444_p3) => function (_445_i2) {
          return _444_p3;
        })(p3))).length), new BigNumber((_443_v17).length)), p3));
        let _index69 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_442_v16).length));
        (_442_v16)[_index69] = _module.__default.fm15(p3, globalState);
        if (!(p0) || (false)) {
          let _446_v18;
          let _nw77 = Array((new BigNumber(24)).toNumber()).fill(false);
          _446_v18 = _nw77;
          let _index70 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_446_v18).length));
          (_446_v18)[_index70] = p2;
          let _447_v19;
          _447_v19 = _dafny.Seq.UnicodeFromString("ixhgw");
          let _index71 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_446_v18).length));
          (_446_v18)[_index71] = _dafny.areEqual(_447_v19, _dafny.Seq.Concat(_447_v19, _dafny.Seq.UnicodeFromString("md")));
          let _448_v20;
          _448_v20 = _module.D2.create_DC4();
          let _449_v21;
          _449_v21 = _dafny.MultiSet.fromElements(_448_v20);
          _449_v21 = _449_v21;
          r2 = p2;
          let _450_v22;
          _450_v22 = _dafny.MultiSet.fromElements(_443_v17);
          let _451_v23;
          _451_v23 = _dafny.Seq.of((_442_v16)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_442_v16).length))], (_442_v16)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_442_v16).length))]);
          r2 = (_450_v22).IsProperSubsetOf((_dafny.MultiSet.FromArray(_451_v23)).Difference(_dafny.MultiSet.fromElements((_442_v16)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_442_v16).length))])));
          let _452_v24;
          let _nw78 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _452_v24 = _nw78;
          let _453_v25;
          _453_v25 = _dafny.Seq.of(_452_v24, _452_v24);
          let _454_v26;
          _454_v26 = _module.D6.create_DC12(_453_v25);
          let _455_v27;
          _455_v27 = _module.D6.create_DC12((_454_v26).dtor_cf17);
          let _456_v28;
          _456_v28 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm9(new BigNumber(((_455_v27).dtor_cf17).length), globalState)).length));
          _456_v28 = _456_v28;
        } else {
          let _457_v29;
          _457_v29 = _module.D1.create_DC2();
          let _458_v30;
          let _459_v31;
          let _460_v32;
          let _461_v33;
          let _out20;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector6 = (_this).m3(_module.__default.fm16(p1, p3, p2, globalState), (_dafny.ZERO).minus(p3), (p2) && ((_this).f18), _457_v29, globalState);
          _out20 = _outcollector6[0];
          _out21 = _outcollector6[1];
          _out22 = _outcollector6[2];
          _out23 = _outcollector6[3];
          _458_v30 = _out20;
          _459_v31 = _out21;
          _460_v32 = _out22;
          _461_v33 = _out23;
          _461_v33 = p3;
          let _462_v34;
          let _nw79 = Array((new BigNumber(4)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _462_v34 = _nw79;
          let _rhs58 = _462_v34;
          let _rhs59 = p0;
          let _rhs60 = !(!(!((p1) && (p1)) || (p1)));
          let _lhs49 = globalState;
          _462_v34 = _rhs58;
          r1 = _rhs59;
          _lhs49.f7 = _rhs60;
          r3 = (_this).f19;
          let _463_v35;
          _463_v35 = _dafny.Seq.UnicodeFromString("ivvmb");
          let _464_v36;
          _464_v36 = _dafny.Map.Empty.slice().updateUnsafe(_463_v35,(_this).f19);
          _464_v36 = (_464_v36).update(_463_v35, (_this).f19);
        }
        let _source10 = _module.D4.create_DC9();
        if (_source10.is_DC9) {
          let _465_v37;
          let _nw80 = new _module.C1();
          _nw80.__ctor((_this).f18, p2, p0, p3);
          _465_v37 = _nw80;
          let _466_v38;
          _466_v38 = _dafny.Seq.UnicodeFromString("gsqmtnha");
          _466_v38 = _466_v38;
          let _467_v39;
          _467_v39 = _dafny.Seq.of((_this).f18, false, false, true, !((_465_v37).f23));
          let _468_v40;
          _468_v40 = _dafny.Seq.of(_467_v39);
          let _469_v41;
          let _nw81 = Array((new BigNumber(24)).toNumber()).fill(false);
          _469_v41 = _nw81;
          let _470_v42;
          _470_v42 = _dafny.Map.Empty.slice().updateUnsafe(_469_v41,p3);
          let _471_v44;
          _471_v44 = _dafny.Seq.of(_module.__default.fm17(_465_v37.f22, (_this).f19, p0, globalState));
          let _472_v45;
          _472_v45 = _module.D2.create_DC3(_471_v44);
          let _473_v46;
          _473_v46 = _module.D2.create_DC5(_472_v45);
          let _474_v47;
          _474_v47 = _dafny.MultiSet.fromElements(_473_v46);
          let _475_v48;
          _475_v48 = _dafny.Map.Empty.slice().updateUnsafe(_473_v46,_dafny.MultiSet.FromArray(_467_v39));
          let _476_v49;
          _476_v49 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
          let _477_v50;
          _477_v50 = _dafny.Set.fromElements(true);
          let _478_v51;
          _478_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_466_v38);
          let _479_v52;
          let _nw82 = Array((new BigNumber(17)).toNumber());
          _nw82[(_dafny.ZERO).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_469_v41,p3)).Merge(_470_v42)).length);
          _nw82[(_dafny.ONE).toNumber()] = (_this).f19;
          _nw82[(new BigNumber(2)).toNumber()] = ((_this).f19).multipliedBy(new BigNumber((_466_v38).length));
          _nw82[(new BigNumber(3)).toNumber()] = p3;
          _nw82[(new BigNumber(4)).toNumber()] = (_this).f19;
          _nw82[(new BigNumber(5)).toNumber()] = new BigNumber(((function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of (_474_v47).Elements) {
              let _480_v43 = _compr_16;
              if ((_474_v47).contains(_480_v43)) {
                _coll16.push([_480_v43,_dafny.MultiSet.fromElements(_465_v37.f22)]);
              }
            }
            return _coll16;
          }()).Merge(_475_v48)).length);
          _nw82[(new BigNumber(6)).toNumber()] = p3;
          _nw82[(new BigNumber(7)).toNumber()] = new BigNumber((_476_v49).length);
          _nw82[(new BigNumber(8)).toNumber()] = (_this).f19;
          _nw82[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_467_v39).length), new BigNumber((_477_v50).length));
          _nw82[(new BigNumber(10)).toNumber()] = (p3).plus((_this).f19);
          _nw82[(new BigNumber(11)).toNumber()] = _module.__default.fm0(globalState);
          _nw82[(new BigNumber(12)).toNumber()] = (_this).f19;
          _nw82[(new BigNumber(13)).toNumber()] = p3;
          _nw82[(new BigNumber(14)).toNumber()] = (_this).f19;
          _nw82[(new BigNumber(15)).toNumber()] = new BigNumber((_478_v51).length);
          _nw82[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt((_this).f19, new BigNumber(325));
          _479_v52 = _nw82;
          let _index72 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_479_v52).length));
          (_479_v52)[_index72] = (_this).f19;
          let _481_v53;
          _481_v53 = _module.D5.create_DC11(p2, new BigNumber((_443_v17).length), (_this).f18, _466_v38, _module.__default.fm0(globalState));
          let _482_v54;
          _482_v54 = new _dafny.CodePoint('u'.codePointAt(0));
          let _483_v55;
          _483_v55 = _dafny.Map.Empty.slice().updateUnsafe(_479_v52,_dafny.Seq.UnicodeFromString("d"));
          let _pat_let_tv16 = p2;
          let _484_v56;
          let _nw83 = Array((new BigNumber(26)).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = _481_v53;
          _nw83[(_dafny.ONE).toNumber()] = _module.__default.fm18(_481_v53, new _dafny.CodePoint('n'.codePointAt(0)), new BigNumber(((_442_v16)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_442_v16).length))]).length), globalState);
          _nw83[(new BigNumber(2)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(3)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(4)).toNumber()] = ((_465_v37.f22) ? (_481_v53) : (_481_v53));
          _nw83[(new BigNumber(5)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(6)).toNumber()] = _module.__default.fm18(_481_v53, (_this).fm4(!((_this).f18), _module.__default.fm16((_this).f18, (_481_v53).dtor_cf13, p2, globalState), _476_v49, globalState), p3, globalState);
          _nw83[(new BigNumber(7)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(8)).toNumber()] = ((_465_v37.f22) ? (_module.D5.create_DC11(p2, p3, (_465_v37).f23, _466_v38, p3)) : (function (_pat_let4_0) {
            return function (_485_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_486_dt__update_hcf12_h0) {
                  return function (_pat_let6_0) {
                    return function (_487_dt__update_hcf13_h0) {
                      return function (_pat_let7_0) {
                        return function (_488_dt__update_hcf14_h0) {
                          return _module.D5.create_DC11(_486_dt__update_hcf12_h0, _487_dt__update_hcf13_h0, _488_dt__update_hcf14_h0, (_485_dt__update__tmp_h0).dtor_cf15, (_485_dt__update__tmp_h0).dtor_cf16);
                        }(_pat_let7_0);
                      }(false);
                    }(_pat_let6_0);
                  }((_this).f19);
                }(_pat_let5_0);
              }(_pat_let_tv16);
            }(_pat_let4_0);
          }(_module.D5.create_DC11((_467_v39)[_module.__default.safeIndex((_this).f19, new BigNumber((_467_v39).length))], new BigNumber(((_442_v16)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_442_v16).length))]).length), p2, _466_v38, new BigNumber((_443_v17).length)))));
          _nw83[(new BigNumber(9)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(10)).toNumber()] = _module.D5.create_DC11(!(!(p1)), (_this).f19, p1, _466_v38, (_this).f19);
          _nw83[(new BigNumber(11)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(12)).toNumber()] = _module.D5.create_DC11(!(p0), p3, (_this).f18, _466_v38, p3);
          _nw83[(new BigNumber(13)).toNumber()] = function (_pat_let8_0) {
            return function (_489_dt__update__tmp_h1) {
              return function (_pat_let9_0) {
                return function (_490_dt__update_hcf16_h0) {
                  return _module.D5.create_DC11((_489_dt__update__tmp_h1).dtor_cf12, (_489_dt__update__tmp_h1).dtor_cf13, (_489_dt__update__tmp_h1).dtor_cf14, (_489_dt__update__tmp_h1).dtor_cf15, _490_dt__update_hcf16_h0);
                }(_pat_let9_0);
              }((_this).f19);
            }(_pat_let8_0);
          }(_481_v53);
          _nw83[(new BigNumber(14)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(15)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(16)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(17)).toNumber()] = _module.__default.fm18(_481_v53, _482_v54, (_this).f19, globalState);
          _nw83[(new BigNumber(18)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(19)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(20)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(21)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(22)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(23)).toNumber()] = _module.D5.create_DC11((_465_v37).f23, p3, (_465_v37).f23, (((_483_v55).contains(_479_v52)) ? ((_483_v55).get(_479_v52)) : (_466_v38)), (_this).f19);
          _nw83[(new BigNumber(24)).toNumber()] = _481_v53;
          _nw83[(new BigNumber(25)).toNumber()] = _481_v53;
          _484_v56 = _nw83;
          let _pat_let_tv17 = p1;
          let _index73 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_484_v56).length));
          (_484_v56)[_index73] = function (_pat_let10_0) {
            return function (_491_dt__update__tmp_h2) {
              return function (_pat_let11_0) {
                return function (_492_dt__update_hcf14_h1) {
                  return function (_pat_let12_0) {
                    return function (_493_dt__update_hcf16_h1) {
                      return _module.D5.create_DC11((_491_dt__update__tmp_h2).dtor_cf12, (_491_dt__update__tmp_h2).dtor_cf13, _492_dt__update_hcf14_h1, (_491_dt__update__tmp_h2).dtor_cf15, _493_dt__update_hcf16_h1);
                    }(_pat_let12_0);
                  }((_module.D4.create_DC8((_this).f19)).dtor_cf10);
                }(_pat_let11_0);
              }(_pat_let_tv17);
            }(_pat_let10_0);
          }(_481_v53);
          let _index74 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_479_v52).length));
          let _index75 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_484_v56).length));
          let _rhs61 = _468_v40;
          let _rhs62 = _module.__default.safeDivisionInt(p3, p3);
          let _rhs63 = _481_v53;
          let _lhs50 = _479_v52;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_479_v52).length));
          let _lhs52 = _484_v56;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_484_v56).length));
          _468_v40 = _rhs61;
          _lhs50[_lhs51] = _rhs62;
          _lhs52[_lhs53] = _rhs63;
          let _494_v57;
          let _nw84 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
          _494_v57 = _nw84;
          let _495_v58;
          _495_v58 = _dafny.Map.Empty.slice().updateUnsafe(_494_v57,_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), ((_496_v54) => function (_497_i3) {
            return _496_v54;
          })(_482_v54)), _module.__default.safeIndex(new BigNumber(822), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), ((_498_v54) => function (_499_i3) {
            return _498_v54;
          })(_482_v54))).length)), new _dafny.CodePoint('v'.codePointAt(0))), _482_v54));
          _495_v58 = (_495_v58).update((((_465_v37).f23) ? (_494_v57) : (_494_v57)), p1);
        } else {
          let _500___mcc_h0 = (_source10).cf10;
          let _501_cf10 = _500___mcc_h0;
          let _502_v59;
          let _nw85 = new _module.C0();
          _nw85.__ctor(true, (_this).f18, (_501_cf10).minus(new BigNumber(458)));
          _502_v59 = _nw85;
          let _503_v60;
          let _init8 = ((_504_cf10, _505_p2) => function (_506_i4) {
            return !(_504_cf10).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_505_p2,new BigNumber(117))).length));
          })(_501_cf10, p2);
          let _nw86 = Array((new BigNumber(20)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw86.length); _i0_8++) {
            _nw86[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _503_v60 = _nw86;
          _503_v60 = _503_v60;
          r3 = _module.__default.safeModuloInt((_dafny.ZERO).minus((((_this).f18) ? ((_this).f19) : ((_dafny.ZERO).minus(_501_cf10)))), _501_cf10);
          let _507_v61;
          let _nw87 = new _module.C0();
          _nw87.__ctor(!(p1), ((!(p2)) ? (p0) : (!(_module.__default.fm16(p2, p3, (_502_v59).f24, globalState)))), p3);
          _507_v61 = _nw87;
          let _508_v62;
          _508_v62 = new _dafny.CodePoint('t'.codePointAt(0));
          let _rhs64 = _507_v61;
          let _rhs65 = _508_v62;
          _507_v61 = _rhs64;
          r0 = _rhs65;
        }
      } else {
        let _509_v63;
        _509_v63 = new _dafny.CodePoint('l'.codePointAt(0));
        let _510_v64;
        _510_v64 = _dafny.Set.fromElements(_509_v63);
        let _511_v65;
        _511_v65 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
        let _512_v66;
        _512_v66 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(((p0) ? (new BigNumber((_dafny.Set.fromElements((_this).f19, new BigNumber((_510_v64).length), (_this).f19, (_this).f19)).length)) : (new BigNumber(-551)))),_511_v65);
        _512_v66 = (_512_v66).update((_this).f19, _511_v65);
        if (p0) {
          r2 = (_this).f18;
          let _513_v67;
          let _nw88 = Array((new BigNumber(25)).toNumber()).fill(_module.D2.Default());
          _513_v67 = _nw88;
          let _514_v68;
          _514_v68 = _module.D2.create_DC4();
          let _index76 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_513_v67).length));
          (_513_v67)[_index76] = _514_v68;
          let _index77 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_513_v67).length));
          let _rhs66 = _module.D2.create_DC4();
          let _rhs67 = _module.__default.fm0(globalState);
          let _lhs54 = _513_v67;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_513_v67).length));
          _lhs54[_lhs55] = _rhs66;
          r3 = _rhs67;
          let _515_v69;
          _515_v69 = _dafny.Seq.of((_this).f19, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,p1)).length));
          let _516_v70;
          _516_v70 = _module.D5.create_DC11(p0, (_this).f19, p2, _dafny.Seq.UnicodeFromString("lmcptqp"), (_this).f19);
          let _rhs68 = new BigNumber((_dafny.Seq.Concat(_515_v69, _dafny.Seq.of((_this).f19))).length);
          let _rhs69 = p1;
          let _rhs70 = ((_516_v70).dtor_cf13).minus(p3);
          let _rhs71 = p3;
          let _lhs56 = globalState;
          r3 = _rhs68;
          _lhs56.f7 = _rhs69;
          r3 = _rhs70;
          r3 = _rhs71;
          r3 = _module.__default.safeModuloInt((new BigNumber(737)).multipliedBy((_this).f19), p3);
          let _517_v71;
          _517_v71 = _module.D6.create_DC13((_this).f19);
          let _518_v72;
          let _nw89 = new _module.C1();
          _nw89.__ctor(p2, p1, !((_this).f18), p3);
          _518_v72 = _nw89;
          let _519_v73;
          _519_v73 = _dafny.Map.Empty.slice().updateUnsafe(_517_v71,_518_v72);
          let _520_v74;
          let _nw90 = Array((new BigNumber(6)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _518_v72;
          _nw90[(_dafny.ONE).toNumber()] = _518_v72;
          _nw90[(new BigNumber(2)).toNumber()] = _518_v72;
          _nw90[(new BigNumber(3)).toNumber()] = _518_v72;
          _nw90[(new BigNumber(4)).toNumber()] = _518_v72;
          _nw90[(new BigNumber(5)).toNumber()] = _518_v72;
          _520_v74 = _nw90;
          let _521_v75;
          _521_v75 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_519_v73).length),_520_v74);
          _521_v75 = (_521_v75).update(new BigNumber(926), _520_v74);
        } else {
          let _522_v76;
          let _nw91 = new _module.C1();
          _nw91.__ctor(p1, ((_dafny.ZERO).minus((_this).f19)).isLessThan((_this).f19), p2, _module.__default.fm0(globalState));
          _522_v76 = _nw91;
          _522_v76 = _522_v76;
          let _523_v77;
          _523_v77 = _module.D6.create_DC13((_this).f19);
          r3 = (p3).plus((_523_v77).dtor_cf18);
          let _524_v78;
          _524_v78 = _dafny.Seq.UnicodeFromString("b");
          let _525_v79;
          let _nw92 = new _module.C0();
          _nw92.__ctor((_522_v76).f23, (((_522_v76).f23) ? (!((_this).f18)) : ((_522_v76).f23)), new BigNumber((_dafny.Seq.Concat(_524_v78, _524_v78)).length));
          _525_v79 = _nw92;
          _509_v63 = _509_v63;
          let _526_v80;
          _526_v80 = _dafny.Seq.of((_525_v79).f24, (_522_v76).f23);
          let _527_v81;
          _527_v81 = _dafny.Seq.of((_522_v76).f23, (_526_v80)[_module.__default.safeIndex((_this).f19, new BigNumber((_526_v80).length))], !(_522_v76.f22) || ((_522_v76).f23), (_525_v79).f24);
          _527_v81 = _527_v81;
        }
        let _528_v82;
        let _init9 = ((_529_p2, _530_p0, _531_p1) => function (_532_i5) {
          return (_dafny.Set.fromElements((_this).f18, _529_p2, _530_p0)).IsDisjointFrom(_dafny.Set.fromElements((_this).f18, _531_p1, (_this).f18, _530_p0));
        })(p2, p0, p1);
        let _nw93 = Array((new BigNumber(21)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw93.length); _i0_9++) {
          _nw93[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _528_v82 = _nw93;
        let _index78 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length));
        (_528_v82)[_index78] = !((_this).f19).isEqualTo((_this).f19);
        let _533_v83;
        let _init10 = ((_534_p3) => function (_535_i6) {
          return (_535_i6).plus(_534_p3);
        })(p3);
        let _nw94 = Array((new BigNumber(15)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw94.length); _i0_10++) {
          _nw94[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _533_v83 = _nw94;
        let _536_v84;
        _536_v84 = _dafny.Map.Empty.slice().updateUnsafe(p0,_533_v83);
        let _537_v85;
        _537_v85 = _module.D1.create_DC1((((_536_v84).contains((_this).f18)) ? ((_536_v84).get((_this).f18)) : (_533_v83)));
        let _538_v86;
        _538_v86 = _dafny.Seq.UnicodeFromString("xqo");
        let _539_v87;
        _539_v87 = _dafny.Set.fromElements((_this).f19);
        let _pat_let_tv18 = _533_v83;
        let _index79 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length));
        let _rhs72 = (_dafny.ZERO).minus(new BigNumber((_539_v87).length));
        let _rhs73 = _dafny.areEqual(_dafny.Seq.Concat(_538_v86, _538_v86), _dafny.Seq.UnicodeFromString("gonr"));
        let _rhs74 = function (_pat_let13_0) {
          return function (_540_dt__update__tmp_h3) {
            return function (_pat_let14_0) {
              return function (_541_dt__update_hcf1_h0) {
                return _module.D1.create_DC1(_541_dt__update_hcf1_h0);
              }(_pat_let14_0);
            }(_pat_let_tv18);
          }(_pat_let13_0);
        }(_537_v85);
        let _rhs75 = _538_v86;
        let _rhs76 = (((_this).f18) ? (p3) : ((_this).f19));
        let _lhs57 = _528_v82;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length));
        r3 = _rhs72;
        _lhs57[_lhs58] = _rhs73;
        _537_v85 = _rhs74;
        _538_v86 = _rhs75;
        r3 = _rhs76;
        if (p1) {
          let _index80 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length));
          (_528_v82)[_index80] = p2;
          let _542_v88;
          _542_v88 = _dafny.Seq.of(_528_v82);
          let _543_v89;
          _543_v89 = _dafny.Seq.of((_this).f18);
          let _544_v90;
          _544_v90 = _dafny.MultiSet.fromElements(_528_v82, _528_v82);
          (globalState).f7 = (_544_v90).IsSubsetOf((((_this).f18) ? (_dafny.MultiSet.fromElements((_542_v88)[_module.__default.safeIndex(new BigNumber((_543_v89).length), new BigNumber((_542_v88).length))], _528_v82, _528_v82, _528_v82)) : (_544_v90)));
          (globalState).f7 = _module.__default.fm16(!((_528_v82)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length))]) || ((_this).f18), p3, p1, globalState);
          let _545_v91;
          _545_v91 = _dafny.Seq.of((_this).f19, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_546_p3) => function (_547_i7) {
            return _546_p3;
          })(p3))).length), (_this).f19);
          let _548_v92;
          _548_v92 = _dafny.Map.Empty.slice().updateUnsafe(_545_v91,_533_v83);
          let _549_v93;
          _549_v93 = _module.D4.create_DC8(p3);
          let _550_v94;
          _550_v94 = _dafny.Map.Empty.slice().updateUnsafe(_549_v93,_dafny.Seq.Concat(_545_v91, _module.__default.fm15(p3, globalState)));
          let _551_v95;
          _551_v95 = _dafny.Map.Empty.slice().updateUnsafe(!((_528_v82)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length))]),(_dafny.ZERO).minus((_this).f19));
          let _552_v96;
          _552_v96 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)));
          let _index81 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length));
          let _rhs77 = (((_548_v92).contains(_module.__default.fm15(p3, globalState))) ? ((_548_v92).get(_module.__default.fm15(p3, globalState))) : (_533_v83));
          let _rhs78 = (_this).f18;
          let _rhs79 = (((_550_v94).contains(_549_v93)) ? ((_550_v94).get(_549_v93)) : (_module.__default.fm15(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("qgvjst"), _module.__default.safeIndex((_this).f19, new BigNumber((_dafny.Seq.UnicodeFromString("qgvjst")).length)), new _dafny.CodePoint('k'.codePointAt(0)))).length), globalState)));
          let _rhs80 = !(!(p2));
          let _rhs81 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_551_v95).length)), (((_552_v96).contains(_509_v63)) ? ((_552_v96).get(_509_v63)) : ((_this).f19)));
          let _lhs59 = globalState;
          let _lhs60 = _528_v82;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length));
          _533_v83 = _rhs77;
          _lhs59.f7 = _rhs78;
          _545_v91 = _rhs79;
          _lhs60[_lhs61] = _rhs80;
          r3 = _rhs81;
          let _553_v97;
          _553_v97 = _dafny.Set.fromElements((_528_v82)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length))], p0, (((_511_v65).contains((_this).f19)) ? ((_511_v65).get((_this).f19)) : (p0)));
          let _554_v98;
          _554_v98 = _dafny.Set.fromElements((_553_v97).Intersect(_553_v97));
          _554_v98 = _554_v98;
        } else {
          r3 = p3;
          let _555_v99;
          _555_v99 = _dafny.Set.fromElements((_this).f18);
          let _rhs82 = (_528_v82)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length))];
          let _rhs83 = new BigNumber((_538_v86).length);
          let _rhs84 = (_555_v99).Difference(_555_v99);
          let _rhs85 = false;
          let _lhs62 = globalState;
          let _lhs63 = globalState;
          _lhs62.f7 = _rhs82;
          r3 = _rhs83;
          _555_v99 = _rhs84;
          _lhs63.f7 = _rhs85;
          let _556_v100;
          _556_v100 = _dafny.Seq.of((_528_v82)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length))]);
          let _557_v101;
          let _nw95 = new _module.C0();
          _nw95.__ctor((_528_v82)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length))], false, new BigNumber((_556_v100).length));
          _557_v101 = _nw95;
          let _558_v102;
          _558_v102 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Seq.UnicodeFromString("utnbh")).length));
          let _559_v103;
          _559_v103 = _module.D3.create_DC7(_557_v101, _558_v102, (_557_v101).f18, _555_v99, p2);
          (globalState).f7 = !((_559_v103).dtor_cf9);
          r1 = p2;
          let _index82 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_533_v83).length));
          (_533_v83)[_index82] = new BigNumber(23);
          let _560_v104;
          _560_v104 = _dafny.MultiSet.fromElements(p3, p3);
          let _index83 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_533_v83).length));
          (_533_v83)[_index83] = (_dafny.ZERO).minus(new BigNumber((_560_v104).cardinality()));
          let _index84 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_533_v83).length));
          let _index85 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_533_v83).length));
          let _rhs86 = _509_v63;
          let _rhs87 = ((_this).f19).plus((new BigNumber(-112)).minus(p3));
          let _rhs88 = p3;
          let _lhs64 = _533_v83;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_533_v83).length));
          let _lhs66 = _533_v83;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_533_v83).length));
          _509_v63 = _rhs86;
          _lhs64[_lhs65] = _rhs87;
          _lhs66[_lhs67] = _rhs88;
        }
        r1 = (_528_v82)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_528_v82).length))];
      }
      let _561_v105;
      _561_v105 = _dafny.Seq.of(p0, p2);
      let _562_v106;
      _562_v106 = _dafny.MultiSet.fromElements(p1, p0, p0);
      let _563_v107;
      _563_v107 = _dafny.Seq.UnicodeFromString("knuyas");
      let _564_v108;
      let _nw96 = Array((new BigNumber(23)).toNumber()).fill(false);
      _564_v108 = _nw96;
      let _565_v109;
      _565_v109 = _dafny.Map.Empty.slice().updateUnsafe(!(!((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,(_561_v105)[_module.__default.safeIndex(p3, new BigNumber((_561_v105).length))])).length))).isEqualTo((((_562_v106).contains(p2)) ? ((_562_v106).get(p2)) : (new BigNumber((_563_v107).length))))),_564_v108);
      let _566_v110;
      _566_v110 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC2(),p2);
      let _567_v112;
      _567_v112 = _module.D1.create_DC2();
      let _568_v113;
      _568_v113 = _dafny.Set.fromElements(_567_v112, _567_v112, _567_v112);
      let _569_v114;
      let _nw97 = Array((new BigNumber(4)).toNumber());
      _nw97[(_dafny.ZERO).toNumber()] = _566_v110;
      _nw97[(_dafny.ONE).toNumber()] = ((p0) ? (_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC2(),(_this).f18)) : (function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_568_v113).Elements) {
          let _570_v111 = _compr_17;
          if ((_568_v113).contains(_570_v111)) {
            _coll17.push([_570_v111,p1]);
          }
        }
        return _coll17;
      }()));
      _nw97[(new BigNumber(2)).toNumber()] = (_566_v110).Merge(_566_v110);
      _nw97[(new BigNumber(3)).toNumber()] = (_566_v110).Merge(_566_v110);
      _569_v114 = _nw97;
      let _index86 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_569_v114).length));
      (_569_v114)[_index86] = (_566_v110).update(_567_v112, p1);
      let _index87 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_569_v114).length));
      let _rhs89 = (((p2) === (p1)) ? (p3) : (_module.__default.fm0(globalState)));
      let _rhs90 = (_565_v109).update((_this).f18, _564_v108);
      let _rhs91 = _566_v110;
      let _lhs68 = _569_v114;
      let _lhs69 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_569_v114).length));
      r3 = _rhs89;
      _565_v109 = _rhs90;
      _lhs68[_lhs69] = _rhs91;
      let _571_v115;
      _571_v115 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),_564_v108);
      let _572_v116;
      _572_v116 = _dafny.MultiSet.fromElements(new BigNumber((_563_v107).length), (_this).f19);
      let _573_v117;
      _573_v117 = _module.D5.create_DC11(p1, (_this).f19, p0, _563_v107, new BigNumber((_572_v116).cardinality()));
      let _574_v118;
      _574_v118 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f19);
      let _575_v119;
      _575_v119 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber(492));
      let _576_v120;
      _576_v120 = _dafny.Set.fromElements((_this).f18);
      let _577_v121;
      let _nw98 = Array((new BigNumber(26)).toNumber());
      _nw98[(_dafny.ZERO).toNumber()] = (_this).f19;
      _nw98[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm0(globalState));
      _nw98[(new BigNumber(2)).toNumber()] = (_573_v117).dtor_cf16;
      _nw98[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f19),_module.__default.fm16(p2, (_dafny.ZERO).minus(p3), false, globalState))).length);
      _nw98[(new BigNumber(4)).toNumber()] = p3;
      _nw98[(new BigNumber(5)).toNumber()] = p3;
      _nw98[(new BigNumber(6)).toNumber()] = p3;
      _nw98[(new BigNumber(7)).toNumber()] = p3;
      _nw98[(new BigNumber(8)).toNumber()] = p3;
      _nw98[(new BigNumber(9)).toNumber()] = (_this).f19;
      _nw98[(new BigNumber(10)).toNumber()] = new BigNumber(((_574_v118).update(p2, new BigNumber((_575_v119).length))).length);
      _nw98[(new BigNumber(11)).toNumber()] = p3;
      _nw98[(new BigNumber(12)).toNumber()] = (_this).f19;
      _nw98[(new BigNumber(13)).toNumber()] = p3;
      _nw98[(new BigNumber(14)).toNumber()] = (_this).f19;
      _nw98[(new BigNumber(15)).toNumber()] = p3;
      _nw98[(new BigNumber(16)).toNumber()] = (_this).f19;
      _nw98[(new BigNumber(17)).toNumber()] = p3;
      _nw98[(new BigNumber(18)).toNumber()] = (_this).f19;
      _nw98[(new BigNumber(19)).toNumber()] = (_this).f19;
      _nw98[(new BigNumber(20)).toNumber()] = new BigNumber((_module.__default.fm9((_this).f19, globalState)).length);
      _nw98[(new BigNumber(21)).toNumber()] = p3;
      _nw98[(new BigNumber(22)).toNumber()] = new BigNumber((_576_v120).length);
      _nw98[(new BigNumber(23)).toNumber()] = (_this).f19;
      _nw98[(new BigNumber(24)).toNumber()] = p3;
      _nw98[(new BigNumber(25)).toNumber()] = (_this).f19;
      _577_v121 = _nw98;
      let _578_v122;
      _578_v122 = _dafny.Map.Empty.slice().updateUnsafe(_577_v121,_564_v108);
      let _579_v123;
      _579_v123 = _module.D3.create_DC6(_564_v108);
      let _580_v124;
      let _nw99 = Array((new BigNumber(24)).toNumber());
      _nw99[(_dafny.ZERO).toNumber()] = _564_v108;
      _nw99[(_dafny.ONE).toNumber()] = _564_v108;
      _nw99[(new BigNumber(2)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(3)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(4)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(5)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(6)).toNumber()] = (((_571_v115).contains(p3)) ? ((_571_v115).get(p3)) : (_564_v108));
      _nw99[(new BigNumber(7)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(8)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(9)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(10)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(11)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(12)).toNumber()] = (((_578_v122).contains(_577_v121)) ? ((_578_v122).get(_577_v121)) : (_564_v108));
      _nw99[(new BigNumber(13)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(14)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(15)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(16)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(17)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(18)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(19)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(20)).toNumber()] = (_579_v123).dtor_cf4;
      _nw99[(new BigNumber(21)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(22)).toNumber()] = _564_v108;
      _nw99[(new BigNumber(23)).toNumber()] = _564_v108;
      _580_v124 = _nw99;
      _580_v124 = _580_v124;
      _564_v108 = _564_v108;
      let _581_v125;
      _581_v125 = new _dafny.CodePoint('c'.codePointAt(0));
      r0 = _581_v125;
      r1 = p0;
      r2 = false;
      r3 = p3;
      return [r0, r1, r2, r3];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = new _dafny.CodePoint('D'.codePointAt(0));
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.ZERO;
      let _582_v0;
      _582_v0 = _dafny.Seq.UnicodeFromString("tkbv");
      let _583_v1;
      _583_v1 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),p1);
      let _584_v2;
      _584_v2 = _dafny.Seq.of(_583_v1, _583_v1, _583_v1);
      if (_dafny.areEqual(_module.__default.fm19(new BigNumber(565), p1, new BigNumber((_582_v0).length), _module.D1.create_DC2(), globalState), _584_v2)) {
        let _585_v3;
        _585_v3 = new _dafny.CodePoint('i'.codePointAt(0));
        let _586_v4;
        let _nw100 = new _module.C1();
        _nw100.__ctor(p0, p2, p2, new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('o'.codePointAt(0)), _585_v3, new _dafny.CodePoint('e'.codePointAt(0)))).length));
        _586_v4 = _nw100;
        let _587_v5;
        _587_v5 = _dafny.Seq.of(p1, p1);
        let _588_v6;
        _588_v6 = _dafny.Map.Empty.slice().updateUnsafe(_586_v4,_dafny.Seq.update(_587_v5, _module.__default.safeIndex(new BigNumber(701), new BigNumber((_587_v5).length)), _module.__default.fm0(globalState)));
        _588_v6 = (_588_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(_586_v4,_587_v5));
        r3 = _module.__default.safeModuloInt(p1, ((_dafny.ZERO).minus(new BigNumber((_587_v5).length))).multipliedBy(p1));
        let _589_v7;
        let _init11 = ((_590_v3) => function (_591_i0) {
          return _dafny.Set.fromElements(_590_v3);
        })(_585_v3);
        let _nw101 = Array((new BigNumber(20)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw101.length); _i0_11++) {
          _nw101[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _589_v7 = _nw101;
        let _592_v8;
        _592_v8 = _dafny.Set.fromElements(_585_v3, _585_v3, _585_v3);
        let _index88 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_589_v7).length));
        (_589_v7)[_index88] = _592_v8;
        let _593_v9;
        _593_v9 = _module.D7.create_DC14(_dafny.Set.fromElements(_585_v3));
        let _pat_let_tv19 = _585_v3;
        let _pat_let_tv20 = _585_v3;
        let _index89 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_589_v7).length));
        (_589_v7)[_index89] = (function (_pat_let15_0) {
          return function (_594_dt__update__tmp_h0) {
            return function (_pat_let16_0) {
              return function (_595_dt__update_hcf19_h0) {
                return _module.D7.create_DC14(_595_dt__update_hcf19_h0);
              }(_pat_let16_0);
            }(_dafny.Set.fromElements(_pat_let_tv19, _pat_let_tv20));
          }(_pat_let15_0);
        }(_593_v9)).dtor_cf19;
        (globalState).f7 = (_586_v4).f23;
        let _596_v10;
        let _nw102 = Array((new BigNumber(20)).toNumber());
        _596_v10 = _nw102;
        let _597_v11;
        _597_v11 = _dafny.Map.Empty.slice().updateUnsafe(_586_v4.f22,_586_v4);
        let _index90 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_596_v10).length));
        (_596_v10)[_index90] = (((_597_v11).contains(p0)) ? ((_597_v11).get(p0)) : (_586_v4));
        let _598_v12;
        _598_v12 = _dafny.Map.Empty.slice().updateUnsafe(_586_v4.f22,(_this).f18);
        let _599_v13;
        let _nw103 = Array((new BigNumber(28)).toNumber());
        _nw103[(_dafny.ZERO).toNumber()] = p2;
        _nw103[(_dafny.ONE).toNumber()] = p0;
        _nw103[(new BigNumber(2)).toNumber()] = true;
        _nw103[(new BigNumber(3)).toNumber()] = p2;
        _nw103[(new BigNumber(4)).toNumber()] = (_586_v4).f23;
        _nw103[(new BigNumber(5)).toNumber()] = p0;
        _nw103[(new BigNumber(6)).toNumber()] = p2;
        _nw103[(new BigNumber(7)).toNumber()] = p2;
        _nw103[(new BigNumber(8)).toNumber()] = p2;
        _nw103[(new BigNumber(9)).toNumber()] = (_this).f18;
        _nw103[(new BigNumber(10)).toNumber()] = (_this).f18;
        _nw103[(new BigNumber(11)).toNumber()] = p0;
        _nw103[(new BigNumber(12)).toNumber()] = p0;
        _nw103[(new BigNumber(13)).toNumber()] = p2;
        _nw103[(new BigNumber(14)).toNumber()] = (((_598_v12).contains(p2)) ? ((_598_v12).get(p2)) : (_586_v4.f22));
        _nw103[(new BigNumber(15)).toNumber()] = (_586_v4).f23;
        _nw103[(new BigNumber(16)).toNumber()] = p2;
        _nw103[(new BigNumber(17)).toNumber()] = (_586_v4).f23;
        _nw103[(new BigNumber(18)).toNumber()] = _586_v4.f22;
        _nw103[(new BigNumber(19)).toNumber()] = !((_586_v4).f23);
        _nw103[(new BigNumber(20)).toNumber()] = _586_v4.f22;
        _nw103[(new BigNumber(21)).toNumber()] = (_this).f18;
        _nw103[(new BigNumber(22)).toNumber()] = (_586_v4).f23;
        _nw103[(new BigNumber(23)).toNumber()] = _module.__default.fm16(p0, new BigNumber((_module.__default.fm14((_this).f19, _586_v4.f22, globalState)).length), (_this).f18, globalState);
        _nw103[(new BigNumber(24)).toNumber()] = false;
        _nw103[(new BigNumber(25)).toNumber()] = !(_586_v4.f22);
        _nw103[(new BigNumber(26)).toNumber()] = (_this).f18;
        _nw103[(new BigNumber(27)).toNumber()] = p0;
        _599_v13 = _nw103;
        let _600_v14;
        _600_v14 = _module.D3.create_DC6(_599_v13);
        let _601_v15;
        let _nw104 = Array((new BigNumber(5)).toNumber());
        _nw104[(_dafny.ZERO).toNumber()] = _600_v14;
        _nw104[(_dafny.ONE).toNumber()] = _600_v14;
        _nw104[(new BigNumber(2)).toNumber()] = _module.D3.create_DC6(_599_v13);
        _nw104[(new BigNumber(3)).toNumber()] = _600_v14;
        _nw104[(new BigNumber(4)).toNumber()] = _600_v14;
        _601_v15 = _nw104;
        let _602_v16;
        _602_v16 = _dafny.Seq.of(_601_v15, _601_v15);
        let _index91 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_596_v10).length));
        let _nw105 = new _module.C1();
        _nw105.__ctor((_this).f18, _dafny.Seq.IsPrefixOf(_602_v16, _602_v16), _586_v4.f22, (_this).f19);
        (_596_v10)[_index91] = _nw105;
      } else {
        r3 = _module.__default.fm0(globalState);
        let _603_v17;
        let _nw106 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _603_v17 = _nw106;
        let _604_v18;
        _604_v18 = _dafny.Set.fromElements(_603_v17, _603_v17, _603_v17, _603_v17, _603_v17);
        let _605_v19;
        _605_v19 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_604_v18).length)));
        if (_dafny.Seq.IsPrefixOf(_605_v19, _605_v19)) {
          let _606_v20;
          let _nw107 = new _module.C0();
          _nw107.__ctor(p2, !(p2) || (true), new BigNumber((_582_v0).length));
          _606_v20 = _nw107;
          let _607_v21;
          let _nw108 = new _module.C0();
          _nw108.__ctor((_this).f18, p2, (_this).f19);
          _607_v21 = _nw108;
          let _608_v22;
          let _init12 = ((_609_v21, _610_v20) => function (_611_i1) {
            return !((_609_v21).f24) || ((_610_v20).f24);
          })(_607_v21, _606_v20);
          let _nw109 = Array((new BigNumber(12)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw109.length); _i0_12++) {
            _nw109[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _608_v22 = _nw109;
          r0 = _608_v22;
          let _612_v23;
          _612_v23 = _dafny.Map.Empty.slice().updateUnsafe((_606_v20).f24,p0);
          let _613_v24;
          _613_v24 = _dafny.Seq.of(_module.__default.fm20(globalState), _612_v23, _612_v23, (_dafny.Map.Empty.slice().updateUnsafe((_607_v21).f24,(_606_v20).f24)).Merge(_612_v23), _612_v23);
          let _614_v25;
          _614_v25 = _dafny.Map.Empty.slice().updateUnsafe(_582_v0,p0);
          let _615_v26;
          _615_v26 = new _dafny.CodePoint('e'.codePointAt(0));
          let _616_v27;
          _616_v27 = _dafny.Set.fromElements(_615_v26);
          let _617_v28;
          _617_v28 = _module.D7.create_DC14(_616_v27);
          let _pat_let_tv21 = _615_v26;
          _613_v24 = _module.__default.fm21((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(367)), function (_618_i2) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          }),(_this).f18)).Merge(_614_v25), p1, function (_pat_let17_0) {
            return function (_619_dt__update__tmp_h1) {
              return function (_pat_let18_0) {
                return function (_620_dt__update_hcf19_h1) {
                  return _module.D7.create_DC14(_620_dt__update_hcf19_h1);
                }(_pat_let18_0);
              }(_dafny.Set.fromElements(_pat_let_tv21, new _dafny.CodePoint('a'.codePointAt(0))));
            }(_pat_let17_0);
          }(_617_v28), (new BigNumber(374)).isEqualTo(p1), globalState);
          r1 = _615_v26;
        } else {
          (globalState).f7 = p0;
          let _621_v29;
          _621_v29 = _dafny.Seq.of(p3);
          let _622_v30;
          _622_v30 = _module.D2.create_DC3(_dafny.Seq.Concat(_621_v29, _621_v29));
          let _623_v31;
          let _nw110 = Array((new BigNumber(11)).toNumber()).fill([]);
          _623_v31 = _nw110;
          let _624_v32;
          let _nw111 = Array((new BigNumber(11)).toNumber());
          _nw111[(_dafny.ZERO).toNumber()] = p0;
          _nw111[(_dafny.ONE).toNumber()] = (_this).f18;
          _nw111[(new BigNumber(2)).toNumber()] = (_this).f18;
          _nw111[(new BigNumber(3)).toNumber()] = p0;
          _nw111[(new BigNumber(4)).toNumber()] = p0;
          _nw111[(new BigNumber(5)).toNumber()] = p0;
          _nw111[(new BigNumber(6)).toNumber()] = p0;
          _nw111[(new BigNumber(7)).toNumber()] = (_this).f18;
          _nw111[(new BigNumber(8)).toNumber()] = p2;
          _nw111[(new BigNumber(9)).toNumber()] = (_this).f18;
          _nw111[(new BigNumber(10)).toNumber()] = p0;
          _624_v32 = _nw111;
          let _index92 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_623_v31).length));
          (_623_v31)[_index92] = _624_v32;
          let _625_v33;
          let _nw112 = new _module.C1();
          _nw112.__ctor(p2, p0, p0, new BigNumber(543));
          _625_v33 = _nw112;
          let _626_v34;
          let _nw113 = Array((new BigNumber(10)).toNumber());
          _nw113[(_dafny.ZERO).toNumber()] = _625_v33;
          _nw113[(_dafny.ONE).toNumber()] = _625_v33;
          _nw113[(new BigNumber(2)).toNumber()] = _625_v33;
          _nw113[(new BigNumber(3)).toNumber()] = _625_v33;
          _nw113[(new BigNumber(4)).toNumber()] = _625_v33;
          _nw113[(new BigNumber(5)).toNumber()] = _625_v33;
          _nw113[(new BigNumber(6)).toNumber()] = _625_v33;
          _nw113[(new BigNumber(7)).toNumber()] = _625_v33;
          _nw113[(new BigNumber(8)).toNumber()] = _625_v33;
          _nw113[(new BigNumber(9)).toNumber()] = _625_v33;
          _626_v34 = _nw113;
          let _index93 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_623_v31).length));
          let _rhs92 = _624_v32;
          let _rhs93 = _622_v30;
          let _rhs94 = _624_v32;
          let _rhs95 = _626_v34;
          let _rhs96 = p0;
          let _lhs70 = _623_v31;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_623_v31).length));
          let _lhs72 = globalState;
          r0 = _rhs92;
          _622_v30 = _rhs93;
          _lhs70[_lhs71] = _rhs94;
          _626_v34 = _rhs95;
          _lhs72.f7 = _rhs96;
          let _627_v35;
          _627_v35 = _dafny.MultiSet.fromElements((_625_v33).f18);
          let _628_v36;
          _628_v36 = new _dafny.CodePoint('d'.codePointAt(0));
          let _629_v37;
          _629_v37 = _dafny.MultiSet.fromElements((_this).fm5(!(p2), _628_v36, p0, (_625_v33).f18, globalState), _627_v35, (_this).fm5((_625_v33).f18, _628_v36, p0, false, globalState), _627_v35);
          let _630_v38;
          _630_v38 = _dafny.Map.Empty.slice().updateUnsafe(_627_v35,(((_629_v37).contains(_dafny.MultiSet.fromElements(p0))) ? ((_629_v37).get(_dafny.MultiSet.fromElements(p0))) : (new BigNumber((_627_v35).cardinality()))));
          _630_v38 = (_630_v38).update(_627_v35, (_625_v33).f19);
          let _631_v39;
          let _nw114 = new _module.C1();
          _nw114.__ctor(!(!((_625_v33).f18)), ((_this).f18) && ((_625_v33).f18), !((_625_v33).f19).isEqualTo(new BigNumber((_582_v0).length)), p1);
          _631_v39 = _nw114;
          (globalState).f7 = (_this).f18;
        }
        let _632_v40;
        _632_v40 = _dafny.MultiSet.fromElements((_this).f19);
        let _633_v41;
        _633_v41 = _dafny.Set.fromElements((_this).f18);
        let _634_v42;
        _634_v42 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
        let _635_v43;
        _635_v43 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(87));
        let _636_v44;
        let _nw115 = Array((new BigNumber(11)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = p2;
        _nw115[(_dafny.ONE).toNumber()] = (_632_v40).contains(p1);
        _nw115[(new BigNumber(2)).toNumber()] = p2;
        _nw115[(new BigNumber(3)).toNumber()] = (_633_v41).IsProperSubsetOf(_633_v41);
        _nw115[(new BigNumber(4)).toNumber()] = !(p0);
        _nw115[(new BigNumber(5)).toNumber()] = ((_this).f19).isLessThanOrEqualTo((_this).f19);
        _nw115[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p2,(((_634_v42).contains(p0)) ? ((_634_v42).get(p0)) : (!(true))))).contains(p0);
        _nw115[(new BigNumber(7)).toNumber()] = !(p0);
        _nw115[(new BigNumber(8)).toNumber()] = true;
        _nw115[(new BigNumber(9)).toNumber()] = (p0) === (p0);
        _nw115[(new BigNumber(10)).toNumber()] = !((_this).f19).isEqualTo(new BigNumber((_635_v43).length));
        _636_v44 = _nw115;
        let _index94 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_636_v44).length));
        (_636_v44)[_index94] = (_this).f18;
        let _637_v45;
        _637_v45 = _dafny.Set.fromElements((new BigNumber((_633_v41).length)).minus(p1));
        let _638_v46;
        let _nw116 = new _module.C1();
        _nw116.__ctor((_this).f18, p0, (_this).f18, p1);
        _638_v46 = _nw116;
        let _639_v47;
        _639_v47 = _dafny.Seq.of(_638_v46);
        let _640_v48;
        _640_v48 = _dafny.Seq.of(_582_v0, _582_v0);
        let _641_v49;
        _641_v49 = _module.D8.create_DC17(_module.__default.fm22(globalState));
        let _index95 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_636_v44).length));
        let _rhs97 = !_dafny.Seq.contains(_dafny.Seq.Concat(_639_v47, _639_v47), _638_v46);
        let _rhs98 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-728)), function (_642_i3) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        })).length)).isLessThanOrEqualTo((new BigNumber(((_640_v48)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_this).f19, new BigNumber(-468))).length)), new BigNumber((_640_v48).length))]).length)).plus(new BigNumber(747)));
        let _rhs99 = (_641_v49).dtor_cf22;
        let _rhs100 = (_638_v46).f23;
        let _lhs73 = globalState;
        let _lhs74 = _636_v44;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_636_v44).length));
        let _lhs76 = globalState;
        _lhs73.f7 = _rhs97;
        _lhs74[_lhs75] = _rhs98;
        _637_v45 = _rhs99;
        _lhs76.f7 = _rhs100;
        let _643_v50;
        let _nw117 = Array((new BigNumber(16)).toNumber()).fill(_module.D1.Default());
        _643_v50 = _nw117;
        let _index96 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_643_v50).length));
        (_643_v50)[_index96] = p3;
        let _index97 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_643_v50).length));
        (_643_v50)[_index97] = p3;
        let _644_v51;
        let _nw118 = Array((_dafny.ONE).toNumber());
        _644_v51 = _nw118;
        let _645_v52;
        let _nw119 = new _module.C0();
        _nw119.__ctor(p2, _638_v46.f22, new BigNumber((_634_v42).length));
        _645_v52 = _nw119;
        let _index98 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_644_v51).length));
        (_644_v51)[_index98] = _645_v52;
        let _index99 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_644_v51).length));
        (_644_v51)[_index99] = _645_v52;
      }
      let _hi2 = _module.__default.safeModuloInt((_this).f19, p1);
      for (let _646_i4 = p1; _646_i4.isLessThan(_hi2); _646_i4 = _646_i4.plus(_dafny.ONE)) {
        let _647_v53;
        let _nw120 = new _module.C0();
        _nw120.__ctor(p2, p0, (_this).f19);
        _647_v53 = _nw120;
        (globalState).f7 = false;
        let _648_v54;
        let _init13 = ((_649_i4) => function (_650_i5) {
          return (_650_i5).plus(_649_i4);
        })(_646_i4);
        let _nw121 = Array((new BigNumber(2)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw121.length); _i0_13++) {
          _nw121[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _648_v54 = _nw121;
        let _index100 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_648_v54).length));
        (_648_v54)[_index100] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f19)).length);
        let _651_v55;
        let _init14 = ((_652_v53) => function (_653_i6) {
          return (((_652_v53).f24) ? (!((_652_v53).f24)) : (true));
        })(_647_v53);
        let _nw122 = Array((new BigNumber(16)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw122.length); _i0_14++) {
          _nw122[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _651_v55 = _nw122;
        let _index101 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_651_v55).length));
        (_651_v55)[_index101] = p0;
        let _654_v56;
        _654_v56 = _dafny.Seq.of(_582_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(361)), function (_655_i7) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }));
        let _index102 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_648_v54).length));
        let _index103 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_651_v55).length));
        let _rhs101 = (_647_v53).f24;
        let _rhs102 = _module.__default.fm0(globalState);
        let _rhs103 = ((_this).f19).plus(p1);
        let _rhs104 = _dafny.Seq.IsProperPrefixOf(_654_v56, _dafny.Seq.update(((p0) ? (_654_v56) : (_654_v56)), _module.__default.safeIndex((_this).f19, new BigNumber((((p0) ? (_654_v56) : (_654_v56))).length)), _582_v0));
        let _rhs105 = _646_i4;
        let _lhs77 = globalState;
        let _lhs78 = _648_v54;
        let _lhs79 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_648_v54).length));
        let _lhs80 = _651_v55;
        let _lhs81 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_651_v55).length));
        _lhs77.f7 = _rhs101;
        _lhs78[_lhs79] = _rhs102;
        r3 = _rhs103;
        _lhs80[_lhs81] = _rhs104;
        r3 = _rhs105;
        let _656_v57;
        _656_v57 = new _dafny.CodePoint('a'.codePointAt(0));
        let _657_v58;
        let _nw123 = Array((new BigNumber(27)).toNumber());
        _nw123[(_dafny.ZERO).toNumber()] = _656_v57;
        _nw123[(_dafny.ONE).toNumber()] = _656_v57;
        _nw123[(new BigNumber(2)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(3)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(4)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(5)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(6)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(7)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(8)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
        _nw123[(new BigNumber(10)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
        _nw123[(new BigNumber(12)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(13)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(14)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(15)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(16)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(17)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(18)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(19)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(20)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(21)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
        _nw123[(new BigNumber(22)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(23)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(24)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(25)).toNumber()] = _656_v57;
        _nw123[(new BigNumber(26)).toNumber()] = _656_v57;
        _657_v58 = _nw123;
        let _658_v59;
        _658_v59 = _dafny.Map.Empty.slice().updateUnsafe((_648_v54)[_module.__default.safeIndex(new BigNumber(540), new BigNumber((_648_v54).length))],_dafny.Set.fromElements(_657_v58));
        let _659_v60;
        _659_v60 = _dafny.Set.fromElements(_657_v58);
        let _660_v61;
        _660_v61 = _dafny.Set.fromElements(new BigNumber((_582_v0).length));
        let _661_v62;
        _661_v62 = _dafny.Map.Empty.slice().updateUnsafe(_648_v54,_659_v60);
        let _662_v63;
        let _nw124 = Array((new BigNumber(28)).toNumber());
        _nw124[(_dafny.ZERO).toNumber()] = (((_658_v59).contains(_646_i4)) ? ((_658_v59).get(_646_i4)) : (_659_v60));
        _nw124[(_dafny.ONE).toNumber()] = _659_v60;
        _nw124[(new BigNumber(2)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(3)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(4)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(5)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(6)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(7)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(8)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(9)).toNumber()] = (_659_v60).Difference(_659_v60);
        _nw124[(new BigNumber(10)).toNumber()] = (_659_v60).Union(_659_v60);
        _nw124[(new BigNumber(11)).toNumber()] = (_659_v60).Difference(_dafny.Set.fromElements(_657_v58));
        _nw124[(new BigNumber(12)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(13)).toNumber()] = (((_658_v59).contains(p1)) ? ((_658_v59).get(p1)) : (_dafny.Set.fromElements(_657_v58)));
        _nw124[(new BigNumber(14)).toNumber()] = (((_658_v59).contains(new BigNumber((_660_v61).length))) ? ((_658_v59).get(new BigNumber((_660_v61).length))) : (_659_v60));
        _nw124[(new BigNumber(15)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(16)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(17)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(18)).toNumber()] = (_659_v60).Union((((_661_v62).contains(_648_v54)) ? ((_661_v62).get(_648_v54)) : (_659_v60)));
        _nw124[(new BigNumber(19)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(20)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(21)).toNumber()] = (_659_v60).Union(_659_v60);
        _nw124[(new BigNumber(22)).toNumber()] = (_659_v60).Intersect(_659_v60);
        _nw124[(new BigNumber(23)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(24)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(25)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(26)).toNumber()] = _659_v60;
        _nw124[(new BigNumber(27)).toNumber()] = (_659_v60).Union(_659_v60);
        _662_v63 = _nw124;
        let _index104 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_662_v63).length));
        (_662_v63)[_index104] = _659_v60;
        let _index105 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_651_v55).length));
        let _index106 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_662_v63).length));
        let _rhs106 = _656_v57;
        let _rhs107 = _module.__default.safeDivisionInt(_module.__default.fm0(globalState), (_this).f19);
        let _rhs108 = !(!((_this).f18));
        let _rhs109 = (_659_v60).Intersect(_659_v60);
        let _lhs82 = _651_v55;
        let _lhs83 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_651_v55).length));
        let _lhs84 = _662_v63;
        let _lhs85 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_662_v63).length));
        r1 = _rhs106;
        r3 = _rhs107;
        _lhs82[_lhs83] = _rhs108;
        _lhs84[_lhs85] = _rhs109;
      }
      let _663_v64;
      _663_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f18);
      let _664_v65;
      _664_v65 = _dafny.Seq.of(p1, p1, (_this).f19, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length));
      _663_v64 = (_663_v64).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_664_v65).length),(_this).f18));
      if ((_module.__default.fm0(globalState)).isLessThan(p1)) {
        let _665_v66;
        let _nw125 = new _module.C1();
        _nw125.__ctor(p2, true, p0, p1);
        _665_v66 = _nw125;
        let _666_v67;
        _666_v67 = _dafny.Map.Empty.slice().updateUnsafe(p0,_665_v66);
        _666_v67 = (_666_v67).update(p0, _665_v66);
        let _667_v68;
        _667_v68 = _dafny.Seq.of(p2);
        let _668_v69;
        _668_v69 = _dafny.Map.Empty.slice().updateUnsafe(true,_667_v68);
        (globalState).f7 = !(!(_668_v69).contains(false));
        _664_v65 = _dafny.Seq.update(_664_v65, _module.__default.safeIndex(p1, new BigNumber((_664_v65).length)), (_665_v66).f19);
        r3 = p1;
        (globalState).f7 = !(((_this).f19).minus(p1)).isEqualTo(p1);
      } else {
        let _669_v70;
        let _nw126 = Array((new BigNumber(5)).toNumber()).fill(false);
        _669_v70 = _nw126;
        let _index107 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_669_v70).length));
        (_669_v70)[_index107] = (_this).f18;
        let _index108 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_669_v70).length));
        (_669_v70)[_index108] = false;
        let _670_v71;
        let _nw127 = Array((new BigNumber(5)).toNumber());
        _nw127[(_dafny.ZERO).toNumber()] = (_this).f19;
        _nw127[(_dafny.ONE).toNumber()] = new BigNumber(-953);
        _nw127[(new BigNumber(2)).toNumber()] = p1;
        _nw127[(new BigNumber(3)).toNumber()] = (_this).f19;
        _nw127[(new BigNumber(4)).toNumber()] = p1;
        _670_v71 = _nw127;
        let _index109 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_670_v71).length));
        (_670_v71)[_index109] = (_this).f19;
        let _671_v72;
        let _nw128 = new _module.C1();
        _nw128.__ctor(p0, p2, (_this).f18, (_this).f19);
        _671_v72 = _nw128;
        let _672_v73;
        _672_v73 = _dafny.Map.Empty.slice().updateUnsafe(_671_v72,(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,true)).length)).plus(_module.__default.fm0(globalState)));
        let _index110 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_670_v71).length));
        (_670_v71)[_index110] = new BigNumber((_672_v73).length);
        _582_v0 = _582_v0;
        let _673_v74;
        _673_v74 = _module.D2.create_DC4();
        let _674_v75;
        _674_v75 = new _dafny.CodePoint('p'.codePointAt(0));
        (_671_v72).f22 = (_671_v72).fm7(_673_v74, _674_v75, globalState);
        let _675_v76;
        let _nw129 = Array((new BigNumber(15)).toNumber());
        _675_v76 = _nw129;
        let _676_v77;
        let _nw130 = new _module.C0();
        _nw130.__ctor(p0, (_671_v72).f23, (_this).f19);
        _676_v77 = _nw130;
        let _677_v78;
        _677_v78 = _module.D5.create_DC10(_676_v77);
        let _index111 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_675_v76).length));
        (_675_v76)[_index111] = (_677_v78).dtor_cf11;
        let _678_v79;
        let _nw131 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Set.Empty);
        _678_v79 = _nw131;
        let _index112 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_675_v76).length));
        let _rhs110 = _676_v77;
        let _rhs111 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cvhgg"), _582_v0)).length);
        let _rhs112 = _678_v79;
        let _lhs86 = _675_v76;
        let _lhs87 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_675_v76).length));
        _lhs86[_lhs87] = _rhs110;
        r3 = _rhs111;
        _678_v79 = _rhs112;
      }
      (globalState).f7 = !(((_this).f19).isLessThan(new BigNumber(-289))) || ((_this).f18);
      r3 = _module.__default.safeModuloInt((_this).f19, (_this).f19);
      let _679_v80;
      let _init15 = ((_680_p0) => function (_681_i8) {
        return _680_p0;
      })(p0);
      let _nw132 = Array((new BigNumber(8)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw132.length); _i0_15++) {
        _nw132[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _679_v80 = _nw132;
      r0 = _679_v80;
      let _682_v81;
      _682_v81 = _dafny.Set.fromElements((_this).f19);
      let _683_v82;
      _683_v82 = new _dafny.CodePoint('i'.codePointAt(0));
      r1 = (((_682_v81).IsDisjointFrom(_682_v81)) ? (_683_v82) : (_683_v82));
      let _684_v83;
      _684_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f19);
      r2 = _684_v83;
      r3 = ((_this).f19).minus((_this).f19);
      return [r0, r1, r2, r3];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f18 = false;
      this._f19 = _dafny.ZERO;
      this._f20 = [];
      this._f21 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    __ctor(f20, f21, f18, f19) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D1.create_DC2()), _dafny.Seq.of(_module.D1.create_DC2(), _module.D1.create_DC2(), _module.D1.create_DC2(), _module.D1.create_DC2())), (_module.D2.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-666)), function (_685_i0) {
  return _module.D1.create_DC2();
}))).dtor_cf2);
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return ((_this).f19).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(false)).length));
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = [];
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _686_v0;
      _686_v0 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(p0, !(p0))).length));
      let _hi3 = new BigNumber((_686_v0).length);
      for (let _687_i0 = p2; _687_i0.isLessThan(_hi3); _687_i0 = _687_i0.plus(_dafny.ONE)) {
        let _688_v1;
        _688_v1 = new _dafny.CodePoint('j'.codePointAt(0));
        _688_v1 = ((_this).f21)[_module.__default.safeIndex(_687_i0, new BigNumber(((_this).f21).length))];
        r3 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-860)), ((_689_v1) => function (_690_i1) {
          return _689_v1;
        })(_688_v1));
        let _691_v2;
        let _init16 = ((_692_v1) => function (_693_i2) {
          return (_693_i2).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(876)), ((_694_v1) => function (_695_i3) {
            return _694_v1;
          })(_692_v1))).length));
        })(_688_v1);
        let _nw133 = Array((new BigNumber(14)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw133.length); _i0_16++) {
          _nw133[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _691_v2 = _nw133;
        _691_v2 = _691_v2;
        let _696_v3;
        _696_v3 = _dafny.Set.fromElements(p0, p0);
        let _697_v4;
        _697_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(883));
        let _698_v5;
        let _699_v6;
        let _out24;
        let _out25;
        let _outcollector7 = (_this).m1(_696_v3, _module.__default.safeModuloInt((_this).f19, new BigNumber((_dafny.Seq.UnicodeFromString("hsryq")).length)), new BigNumber((_697_v4).length), globalState);
        _out24 = _outcollector7[0];
        _out25 = _outcollector7[1];
        _698_v5 = _out24;
        _699_v6 = _out25;
      }
      let _700_v7;
      _700_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-726),p0);
      _700_v7 = (_700_v7).update(p1, (_this).f18);
      let _701_v8;
      _701_v8 = new _dafny.CodePoint('b'.codePointAt(0));
      let _702_v9;
      _702_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f18);
      let _703_v10;
      let _nw134 = Array((new BigNumber(17)).toNumber());
      _nw134[(_dafny.ZERO).toNumber()] = p0;
      _nw134[(_dafny.ONE).toNumber()] = p0;
      _nw134[(new BigNumber(2)).toNumber()] = (_this).f18;
      _nw134[(new BigNumber(3)).toNumber()] = (p0) || (!((_this).f18));
      _nw134[(new BigNumber(4)).toNumber()] = p0;
      _nw134[(new BigNumber(5)).toNumber()] = p0;
      _nw134[(new BigNumber(6)).toNumber()] = (_this).f18;
      _nw134[(new BigNumber(7)).toNumber()] = p0;
      _nw134[(new BigNumber(8)).toNumber()] = (p1).isLessThanOrEqualTo(p1);
      _nw134[(new BigNumber(9)).toNumber()] = (_this).fm2(_dafny.Seq.update((_this).f21, _module.__default.safeIndex(new BigNumber(((_this).f21).length), new BigNumber(((_this).f21).length)), _701_v8), new BigNumber(((_this).f21).length), globalState);
      _nw134[(new BigNumber(10)).toNumber()] = p0;
      _nw134[(new BigNumber(11)).toNumber()] = (_module.__default.fm0(globalState)).isLessThanOrEqualTo(p1);
      _nw134[(new BigNumber(12)).toNumber()] = !((_this).f18);
      _nw134[(new BigNumber(13)).toNumber()] = p0;
      _nw134[(new BigNumber(14)).toNumber()] = p0;
      _nw134[(new BigNumber(15)).toNumber()] = ((true) ? ((((_700_v7).contains(p1)) ? ((_700_v7).get(p1)) : (p0))) : ((((_702_v9).contains((_this).f21)) ? ((_702_v9).get((_this).f21)) : (!((_this).f18)))));
      _nw134[(new BigNumber(16)).toNumber()] = (_this).fm2((_this).f21, (_this).f19, globalState);
      _703_v10 = _nw134;
      let _704_v11;
      _704_v11 = _module.D3.create_DC6(_703_v10);
      _703_v10 = (_704_v11).dtor_cf4;
      let _705_v12;
      _705_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _hi4 = new BigNumber((_705_v12).length);
      for (let _706_i4 = (_this).f19; _706_i4.isLessThan(_hi4); _706_i4 = _706_i4.plus(_dafny.ONE)) {
        let _707_v13;
        _707_v13 = _dafny.MultiSet.fromElements(p0);
        let _708_v14;
        _708_v14 = _dafny.Seq.of((_this).f19, new BigNumber((_707_v13).cardinality()), p2, p2);
        _708_v14 = _708_v14;
        let _709_v15;
        _709_v15 = new BigNumber(103);
        let _710_v16;
        _710_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(427)),new BigNumber(763));
        let _711_v17;
        let _init17 = ((_712_p2) => function (_713_i5) {
          return (_713_i5).multipliedBy(_712_p2);
        })(p2);
        let _nw135 = Array((new BigNumber(2)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw135.length); _i0_17++) {
          _nw135[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _711_v17 = _nw135;
        let _index113 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_711_v17).length));
        (_711_v17)[_index113] = _706_i4;
        let _714_v19;
        _714_v19 = _dafny.Seq.of(_708_v14, _708_v14, _708_v14);
        let _index114 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_711_v17).length));
        let _rhs113 = _709_v15;
        let _rhs114 = (function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of (_714_v19).Elements) {
            let _715_v18 = _compr_18;
            if (_dafny.Seq.contains(_714_v19, _715_v18)) {
              _coll18.push([_715_v18,_709_v15]);
            }
          }
          return _coll18;
        }()).Merge(_710_v16);
        let _rhs115 = ((_dafny.ZERO).minus((_this).f19)).plus(p2);
        let _rhs116 = _706_i4;
        let _rhs117 = (_this).fm2(_module.__default.fm3(globalState), ((_this).f19).multipliedBy((_dafny.ZERO).minus(_module.__default.fm0(globalState))), globalState);
        let _lhs88 = _711_v17;
        let _lhs89 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_711_v17).length));
        let _lhs90 = globalState;
        _709_v15 = _rhs113;
        _710_v16 = _rhs114;
        _709_v15 = _rhs115;
        _lhs88[_lhs89] = _rhs116;
        _lhs90.f7 = _rhs117;
        if ((_this).f18) {
          _711_v17 = _711_v17;
          _705_v12 = (_705_v12).Merge(_705_v12);
          let _716_v20;
          _716_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,p1);
          let _717_v21;
          _717_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_716_v20);
          let _718_v22;
          _718_v22 = _dafny.MultiSet.fromElements(_706_i4, _706_i4, _module.__default.fm0(globalState), _709_v15, _module.__default.fm0(globalState));
          let _719_v23;
          _719_v23 = _dafny.MultiSet.fromElements(new BigNumber((_717_v21).length), (((_718_v22).contains((_dafny.ZERO).minus(p1))) ? ((_718_v22).get((_dafny.ZERO).minus(p1))) : ((_711_v17)[_module.__default.safeIndex(new BigNumber(419), new BigNumber((_711_v17).length))])), new BigNumber((_705_v12).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_711_v17).length));
          let _rhs118 = !(((_719_v23).Difference((_718_v22).update(new BigNumber(-655), _module.__default.abs(p2)))).IsProperSubsetOf(_719_v23));
          let _rhs119 = _706_i4;
          let _rhs120 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f21, (_this).f21), _dafny.Seq.Concat((_this).f21, (_this).f21));
          let _rhs121 = p1;
          let _rhs122 = !(p0);
          let _lhs91 = globalState;
          let _lhs92 = _711_v17;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_711_v17).length));
          let _lhs94 = globalState;
          _lhs91.f7 = _rhs118;
          _709_v15 = _rhs119;
          r2 = _rhs120;
          _lhs92[_lhs93] = _rhs121;
          _lhs94.f7 = _rhs122;
          let _index116 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_711_v17).length));
          (_711_v17)[_index116] = (_708_v14)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat((_this).f21, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-925)), ((_720_v8) => function (_721_i6) {
            return _720_v8;
          })(_701_v8)))).length), new BigNumber((_708_v14).length))];
          let _722_v24;
          _722_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,_703_v10);
          let _723_v25;
          _723_v25 = _dafny.Seq.of((_this).f20);
          let _724_v26;
          let _nw136 = Array((new BigNumber(29)).toNumber());
          _nw136[(_dafny.ZERO).toNumber()] = (((_722_v24).contains((_this).f18)) ? ((_722_v24).get((_this).f18)) : (_703_v10));
          _nw136[(_dafny.ONE).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(2)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(3)).toNumber()] = (_723_v25)[_module.__default.safeIndex(_706_i4, new BigNumber((_723_v25).length))];
          _nw136[(new BigNumber(4)).toNumber()] = _703_v10;
          _nw136[(new BigNumber(5)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(6)).toNumber()] = (_723_v25)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_723_v25).length))];
          _nw136[(new BigNumber(7)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(8)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(9)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(10)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(11)).toNumber()] = _703_v10;
          _nw136[(new BigNumber(12)).toNumber()] = (_723_v25)[_module.__default.safeIndex(_706_i4, new BigNumber((_723_v25).length))];
          _nw136[(new BigNumber(13)).toNumber()] = _703_v10;
          _nw136[(new BigNumber(14)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(15)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(16)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(17)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(18)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(19)).toNumber()] = _703_v10;
          _nw136[(new BigNumber(20)).toNumber()] = _703_v10;
          _nw136[(new BigNumber(21)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(22)).toNumber()] = (((_722_v24).contains(!((_this).f18))) ? ((_722_v24).get(!((_this).f18))) : ((_this).f20));
          _nw136[(new BigNumber(23)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(24)).toNumber()] = (_this).f20;
          _nw136[(new BigNumber(25)).toNumber()] = _703_v10;
          _nw136[(new BigNumber(26)).toNumber()] = _703_v10;
          _nw136[(new BigNumber(27)).toNumber()] = _703_v10;
          _nw136[(new BigNumber(28)).toNumber()] = (_723_v25)[_module.__default.safeIndex(p2, new BigNumber((_723_v25).length))];
          _724_v26 = _nw136;
          let _index117 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_724_v26).length));
          (_724_v26)[_index117] = (_this).f20;
          let _index118 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_724_v26).length));
          let _nw137 = Array((new BigNumber(15)).toNumber()).fill(false);
          (_724_v26)[_index118] = _nw137;
        } else {
          (globalState).f7 = !(false);
          let _725_v27;
          let _nw138 = new _module.C1();
          _nw138.__ctor(true, !((_this).f18), p0, (_711_v17)[_module.__default.safeIndex(new BigNumber(419), new BigNumber((_711_v17).length))]);
          _725_v27 = _nw138;
          let _726_v28;
          _726_v28 = _dafny.Set.fromElements((_725_v27).f18);
          let _727_v29;
          _727_v29 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC7(_725_v27, _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f19), p0, _726_v28, p0),_dafny.Set.fromElements(_709_v15));
          let _728_v30;
          _728_v30 = _dafny.Seq.of((_this).f18, (_this).fm2((_this).f21, new BigNumber((_727_v29).length), globalState));
          let _729_v31;
          _729_v31 = _module.D5.create_DC11((_725_v27).f18, _706_i4, (_this).f18, (_this).f21, (_this).f19);
          (globalState).f7 = (((_700_v7).contains(new BigNumber((_dafny.Seq.Concat(_728_v30, _dafny.Seq.of((_this).f18, (_729_v31).dtor_cf12))).length))) ? ((_700_v7).get(new BigNumber((_dafny.Seq.Concat(_728_v30, _dafny.Seq.of((_this).f18, (_729_v31).dtor_cf12))).length))) : ((_this).f18));
          let _730_v32;
          let _nw139 = new _module.C0();
          _nw139.__ctor(!((_725_v27).f18), (_725_v27).f18, (_dafny.ZERO).minus(new BigNumber(((_this).f21).length)));
          _730_v32 = _nw139;
          _730_v32 = _730_v32;
          let _731_v33;
          let _nw140 = new _module.C1();
          _nw140.__ctor((_730_v32).f24, false, p0, (_709_v15).plus(_706_i4));
          _731_v33 = _nw140;
        }
        _709_v15 = (_dafny.ZERO).minus(p2);
      }
      let _732_v34;
      _732_v34 = _dafny.Seq.of((_this).f18);
      let _rhs123 = (_this).f21;
      let _rhs124 = ((_module.__default.fm20(globalState)).update(p0, (_this).f18)).update((_dafny.MultiSet.FromArray(_732_v34)).IsSubsetOf(_dafny.MultiSet.FromArray(_732_v34)), (true) && (false));
      let _lhs95 = globalState;
      r2 = _rhs123;
      _lhs95.f6 = _rhs124;
      let _source11 = ((p0) ? (_704_v11) : (_module.D3.create_DC6(_703_v10)));
      if (_source11.is_DC7) {
        let _733___mcc_h0 = (_source11).cf5;
        let _734___mcc_h1 = (_source11).cf6;
        let _735___mcc_h2 = (_source11).cf7;
        let _736___mcc_h3 = (_source11).cf8;
        let _737___mcc_h4 = (_source11).cf9;
        let _738_cf9 = _737___mcc_h4;
        let _739_cf8 = _736___mcc_h3;
        let _740_cf7 = _735___mcc_h2;
        let _741_cf6 = _734___mcc_h1;
        let _742_cf5 = _733___mcc_h0;
        if (_module.__default.fm16(((_this).f19).isLessThan(p1), p2, (_this).f18, globalState)) {
          r0 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("x"), (((_742_cf5).f18) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(636)), ((_743_v8) => function (_744_i7) {
            return _743_v8;
          })(_701_v8))) : ((_this).f21)));
          let _745_v35;
          _745_v35 = _dafny.Map.Empty.slice().updateUnsafe(_701_v8,_703_v10);
          _745_v35 = (_745_v35).update(_701_v8, _703_v10);
          (globalState).f7 = !(_740_cf7);
          let _746_v36;
          let _747_v37;
          let _out26;
          let _out27;
          let _outcollector8 = (_this).m1(_739_cf8, (_module.__default.fm0(globalState)).plus(p2), (_742_cf5).f19, globalState);
          _out26 = _outcollector8[0];
          _out27 = _outcollector8[1];
          _746_v36 = _out26;
          _747_v37 = _out27;
          let _748_v38;
          _748_v38 = _module.D4.create_DC9();
          let _749_v39;
          _749_v39 = _dafny.Map.Empty.slice().updateUnsafe(_748_v38,(_742_cf5).f19);
          _749_v39 = (_749_v39).update(_748_v38, _747_v37);
        } else {
          (globalState).f7 = (((_700_v7).contains((_742_cf5).f19)) ? ((_700_v7).get((_742_cf5).f19)) : ((new BigNumber((_dafny.Seq.of(_738_cf9)).length)).isLessThanOrEqualTo(p1)));
          (globalState).f7 = (_this).f18;
          let _750_v40;
          let _nw141 = new _module.C2();
          _nw141.__ctor(!(_740_cf7), p2);
          _750_v40 = _nw141;
          let _751_v41;
          _751_v41 = _dafny.Map.Empty.slice().updateUnsafe(_750_v40,(_742_cf5).f18);
          _705_v12 = (_705_v12).update(p1, (_dafny.ZERO).minus(((_742_cf5).f19).multipliedBy(new BigNumber((_751_v41).length))));
          let _index119 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_703_v10).length));
          (_703_v10)[_index119] = (_this).f18;
          let _index120 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_703_v10).length));
          let _rhs125 = _dafny.Seq.Concat((_this).f21, (_this).f21);
          let _rhs126 = (_this).f18;
          let _rhs127 = (_this).f21;
          let _lhs96 = _703_v10;
          let _lhs97 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_703_v10).length));
          r0 = _rhs125;
          _lhs96[_lhs97] = _rhs126;
          r2 = _rhs127;
          let _752_v42;
          _752_v42 = new BigNumber(349);
          let _753_v43;
          _753_v43 = _dafny.Seq.of((_742_cf5).f19, (_742_cf5).f19, (_742_cf5).f19, new BigNumber(((_this).f21).length), p1);
          _752_v42 = _module.__default.safeModuloInt((_742_cf5).f19, (_753_v43)[_module.__default.safeIndex(p2, new BigNumber((_753_v43).length))]);
        }
        _738_cf9 = ((_740_cf7) ? (_740_cf7) : ((_this).f18));
        let _754_v44;
        _754_v44 = new BigNumber(170);
        _754_v44 = p2;
        _754_v44 = ((_this).f19).multipliedBy((_742_cf5).f19);
      } else {
        let _755___mcc_h5 = (_source11).cf4;
        let _756_cf4 = _755___mcc_h5;
        let _index121 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_756_cf4).length));
        (_756_cf4)[_index121] = (_this).f18;
        let _index122 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_756_cf4).length));
        (_756_cf4)[_index122] = p0;
        let _757_v45;
        _757_v45 = _dafny.Set.fromElements((_this).f21);
        let _758_v46;
        _758_v46 = new BigNumber(763);
        let _759_v47;
        _759_v47 = _dafny.Set.fromElements((((_702_v9).contains((_this).f21)) ? ((_702_v9).get((_this).f21)) : ((_756_cf4)[_module.__default.safeIndex(new BigNumber(874), new BigNumber((_756_cf4).length))])), (_756_cf4)[_module.__default.safeIndex(new BigNumber(874), new BigNumber((_756_cf4).length))]);
        let _760_v48;
        _760_v48 = _module.D4.create_DC9();
        let _761_v49;
        _761_v49 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f19);
        let _762_v50;
        _762_v50 = _dafny.MultiSet.fromElements(_701_v8, _701_v8);
        let _763_v51;
        let _nw142 = Array((new BigNumber(12)).toNumber());
        _nw142[(_dafny.ZERO).toNumber()] = new BigNumber((_761_v49).length);
        _nw142[(_dafny.ONE).toNumber()] = _758_v46;
        _nw142[(new BigNumber(2)).toNumber()] = (_this).f19;
        _nw142[(new BigNumber(3)).toNumber()] = p2;
        _nw142[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_756_cf4)).length);
        _nw142[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_758_v46));
        _nw142[(new BigNumber(6)).toNumber()] = p2;
        _nw142[(new BigNumber(7)).toNumber()] = _758_v46;
        _nw142[(new BigNumber(8)).toNumber()] = p1;
        _nw142[(new BigNumber(9)).toNumber()] = (((_762_v50).contains(_701_v8)) ? ((_762_v50).get(_701_v8)) : (p1));
        _nw142[(new BigNumber(10)).toNumber()] = (_this).f19;
        _nw142[(new BigNumber(11)).toNumber()] = (_this).f19;
        _763_v51 = _nw142;
        let _764_v52;
        _764_v52 = _dafny.MultiSet.fromElements(_763_v51);
        let _765_v53;
        _765_v53 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_764_v52).cardinality()));
        let _766_v54;
        _766_v54 = _dafny.Map.Empty.slice().updateUnsafe(_760_v48,new BigNumber((_765_v53).length));
        let _767_v55;
        _767_v55 = _dafny.Seq.of(p2);
        let _rhs128 = (((_759_v47).IsSubsetOf(_dafny.Set.fromElements(p0, p0, false))) ? (p0) : (_dafny.Seq.IsPrefixOf(_732_v34, _732_v34)));
        let _rhs129 = (_this).f18;
        let _rhs130 = _757_v45;
        let _rhs131 = new BigNumber(((_766_v54).update(_760_v48, (p2).plus(new BigNumber((_767_v55).length)))).length);
        let _rhs132 = p2;
        let _lhs98 = globalState;
        let _lhs99 = globalState;
        _lhs98.f7 = _rhs128;
        _lhs99.f7 = _rhs129;
        _757_v45 = _rhs130;
        _758_v46 = _rhs131;
        _758_v46 = _rhs132;
        let _768_v56;
        _768_v56 = _dafny.MultiSet.fromElements((_756_cf4)[_module.__default.safeIndex(new BigNumber(874), new BigNumber((_756_cf4).length))], (_756_cf4)[_module.__default.safeIndex(new BigNumber(874), new BigNumber((_756_cf4).length))], (_this).f18, false, p0);
        let _769_v57;
        let _770_v58;
        let _out28;
        let _out29;
        let _outcollector9 = (_this).m1((((_756_cf4)[_module.__default.safeIndex(new BigNumber(874), new BigNumber((_756_cf4).length))]) ? (_759_v47) : (_759_v47)), new BigNumber((_768_v56).cardinality()), p1, globalState);
        _out28 = _outcollector9[0];
        _out29 = _outcollector9[1];
        _769_v57 = _out28;
        _770_v58 = _out29;
        _758_v46 = new BigNumber(47);
      }
      r0 = (_this).f21;
      let _771_v59;
      let _nw143 = Array((new BigNumber(8)).toNumber());
      _nw143[(_dafny.ZERO).toNumber()] = (_this).f21;
      _nw143[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("o");
      _nw143[(new BigNumber(2)).toNumber()] = (_this).f21;
      _nw143[(new BigNumber(3)).toNumber()] = (_this).f21;
      _nw143[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(917)), ((_772_v8) => function (_773_i8) {
        return _772_v8;
      })(_701_v8));
      _nw143[(new BigNumber(5)).toNumber()] = (_this).f21;
      _nw143[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((_this).f21, (_this).f21);
      _nw143[(new BigNumber(7)).toNumber()] = (((_this).f18) ? ((_this).f21) : ((_this).f21));
      _771_v59 = _nw143;
      r1 = _771_v59;
      r2 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f21, _dafny.Seq.UnicodeFromString("dgpikn")), (_this).f21);
      r3 = _dafny.Seq.Concat((_this).f21, (_this).f21);
      return [r0, r1, r2, r3];
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _774_v0;
      let _nw144 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
      _774_v0 = _nw144;
      let _775_v1;
      _775_v1 = _dafny.Seq.of((_this).f18);
      let _index123 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_774_v0).length));
      (_774_v0)[_index123] = _775_v1;
      let _index124 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_774_v0).length));
      (_774_v0)[_index124] = _775_v1;
      let _776_v2;
      _776_v2 = _dafny.MultiSet.fromElements(p2);
      let _777_v3;
      _777_v3 = _dafny.MultiSet.fromElements(_776_v2);
      let _778_v4;
      _778_v4 = _dafny.Set.fromElements((((_777_v3).contains(_776_v2)) ? ((_777_v3).get(_776_v2)) : ((_this).f19)), (_dafny.ZERO).minus(p1), new BigNumber(((_this).f21).length), (_dafny.ZERO).minus((_this).f19), p1);
      let _779_i0;
      _779_i0 = _dafny.ZERO;
      L4: {
        while (!(_778_v4).contains(_module.__default.safeModuloInt(new BigNumber(((_774_v0)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_774_v0).length))]).length), p2))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_779_i0)) {
              break L4;
            }
            _779_i0 = (_779_i0).plus(_dafny.ONE);
            let _index125 = _module.__default.safeIndex(new BigNumber(163), new BigNumber(((_this).f20).length));
            ((_this).f20)[_index125] = (_this).f18;
            let _index126 = _module.__default.safeIndex(new BigNumber(163), new BigNumber(((_this).f20).length));
            ((_this).f20)[_index126] = _module.__default.fm16((_this).f18, p1, (new BigNumber(797)).isLessThan(p1), globalState);
            let _index127 = _module.__default.safeIndex(new BigNumber(163), new BigNumber(((_this).f20).length));
            ((_this).f20)[_index127] = (_this).f18;
            let _780_v5;
            let _nw145 = new _module.C0();
            _nw145.__ctor(!_dafny.areEqual(_dafny.Seq.of(!(((_this).f20)[_module.__default.safeIndex(new BigNumber(163), new BigNumber(((_this).f20).length))]), _module.__default.fm16(true, p1, ((_this).f20)[_module.__default.safeIndex(new BigNumber(163), new BigNumber(((_this).f20).length))], globalState)), _775_v1), (_this).f18, p1);
            _780_v5 = _nw145;
            let _781_v6;
            let _init18 = function (_782_i1) {
              return (_782_i1).multipliedBy((_dafny.ZERO).minus((_this).f19));
            };
            let _nw146 = Array((new BigNumber(24)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw146.length); _i0_18++) {
              _nw146[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _781_v6 = _nw146;
            let _index128 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_781_v6).length));
            (_781_v6)[_index128] = p1;
            let _index129 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_781_v6).length));
            (_781_v6)[_index129] = (_this).f19;
          }
        }
      }
      let _783_v7;
      _783_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_module.__default.safeDivisionInt(p1, (_this).f19));
      let _784_v8;
      _784_v8 = new _dafny.CodePoint('n'.codePointAt(0));
      let _785_v9;
      _785_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f21).length),new BigNumber(-598));
      let _rhs133 = (_775_v1)[_module.__default.safeIndex(_module.__default.safeDivisionInt(p2, p2), new BigNumber((_775_v1).length))];
      let _rhs134 = _module.__default.fm16(!((_this).f18), new BigNumber((_dafny.Seq.UnicodeFromString("afvvwktda")).length), (_this).f18, globalState);
      let _rhs135 = (function () {
        let _coll19 = new _dafny.Set();
        for (const _compr_19 of (_785_v9).Keys.Elements) {
          let _786_v10 = _compr_19;
          if ((_785_v9).contains(_786_v10)) {
            _coll19.add((_786_v10).minus(new BigNumber((_dafny.Seq.UnicodeFromString("kuijnymnx")).length)));
          }
        }
        return _coll19;
      }()).IsProperSubsetOf((_dafny.Set.fromElements(new BigNumber((_module.__default.fm23(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,p2), globalState)).length), p1)).Union(_778_v4));
      let _rhs136 = _783_v7;
      let _rhs137 = new _dafny.CodePoint('g'.codePointAt(0));
      let _lhs100 = globalState;
      let _lhs101 = globalState;
      _lhs100.f7 = _rhs133;
      r0 = _rhs134;
      _lhs101.f7 = _rhs135;
      _783_v7 = _rhs136;
      _784_v8 = _rhs137;
      if (false) {
        let _787_v11;
        let _nw147 = Array((_dafny.ONE).toNumber());
        _787_v11 = _nw147;
        let _788_v12;
        _788_v12 = _dafny.Seq.of(p2);
        let _789_v13;
        let _nw148 = new _module.C2();
        _nw148.__ctor((_this).f18, new BigNumber((_788_v12).length));
        _789_v13 = _nw148;
        let _index130 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_787_v11).length));
        (_787_v11)[_index130] = _789_v13;
        let _index131 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_787_v11).length));
        (_787_v11)[_index131] = _789_v13;
        let _790_v14;
        _790_v14 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_789_v13).f18);
        _788_v12 = _dafny.Seq.Concat(_788_v12, _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_790_v14).length))));
        let _791_v15;
        _791_v15 = _dafny.Map.Empty.slice().updateUnsafe((_778_v4).Union(_778_v4),(true) || ((_this).f18));
        let _792_v17;
        _792_v17 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18);
        _791_v15 = (function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of (_module.__default.fm24((_789_v13).f19, globalState)).Elements) {
            let _793_v16 = _compr_20;
            if ((_module.__default.fm24((_789_v13).f19, globalState)).contains(_793_v16)) {
              _coll20.push([_793_v16,(_789_v13).f18]);
            }
          }
          return _coll20;
        }()).Merge(_module.__default.fm25((((_792_v17).contains((_this).f18)) ? ((_792_v17).get((_this).f18)) : (new BigNumber(-304))), globalState));
        r1 = (_789_v13).f19;
        let _794_v18;
        _794_v18 = _dafny.Set.fromElements((_this).f21);
        r0 = !((_module.__default.fm13(globalState)).equals(_794_v18)) || (!((_this).f19).isEqualTo(new BigNumber((_788_v12).length)));
      } else {
        let _795_v19;
        _795_v19 = _module.D7.create_DC15((_this).f18);
        let _796_v20;
        _796_v20 = _module.D7.create_DC16(_795_v19);
        let _797_v21;
        _797_v21 = _module.D7.create_DC16(_796_v20);
        let _798_v22;
        _798_v22 = _module.D7.create_DC16(_796_v20);
        let _index132 = _module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length));
        ((_this).f20)[_index132] = !((_this).f18);
        let _799_v23;
        let _nw149 = Array((new BigNumber(17)).toNumber()).fill(false);
        _799_v23 = _nw149;
        let _800_v24;
        _800_v24 = _dafny.Map.Empty.slice().updateUnsafe(_784_v8,(_this).f18);
        let _index133 = _module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length));
        let _rhs138 = _798_v22;
        let _rhs139 = (_this).f18;
        let _rhs140 = !(((_this).f18) === ((((_800_v24).contains(new _dafny.CodePoint('m'.codePointAt(0)))) ? ((_800_v24).get(new _dafny.CodePoint('m'.codePointAt(0)))) : ((_this).f18))));
        let _rhs141 = !(((_this).f18) === ((_this).f18));
        let _rhs142 = (_this).f20;
        let _lhs102 = (_this).f20;
        let _lhs103 = _module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length));
        let _lhs104 = globalState;
        let _lhs105 = globalState;
        _798_v22 = _rhs138;
        _lhs102[_lhs103] = _rhs139;
        _lhs104.f7 = _rhs140;
        _lhs105.f7 = _rhs141;
        _799_v23 = _rhs142;
        r1 = _module.__default.safeDivisionInt(p2, (_this).f19);
        let _801_v25;
        _801_v25 = _dafny.MultiSet.fromElements((_this).f18);
        let _802_v26;
        _802_v26 = _dafny.Seq.of(new BigNumber((_801_v25).cardinality()));
        r1 = (_802_v26)[_module.__default.safeIndex(new BigNumber(((_774_v0)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_774_v0).length))]).length), new BigNumber((_802_v26).length))];
        let _803_v27;
        let _nw150 = new _module.C2();
        _nw150.__ctor(!((_this).f18), (_this).f19);
        _803_v27 = _nw150;
        let _index134 = _module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length));
        let _rhs143 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_802_v26).length)), p1);
        let _rhs144 = ((_this).f18) && (((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))]);
        let _rhs145 = (_this).f19;
        let _rhs146 = _803_v27;
        let _lhs106 = (_this).f20;
        let _lhs107 = _module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length));
        r1 = _rhs143;
        _lhs106[_lhs107] = _rhs144;
        r1 = _rhs145;
        _803_v27 = _rhs146;
        let _804_v28;
        _804_v28 = _dafny.Set.fromElements(_784_v8);
        if ((_804_v28).IsSubsetOf(_804_v28)) {
          let _index135 = _module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length));
          ((_this).f20)[_index135] = (_this).f18;
          let _index136 = _module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length));
          ((_this).f20)[_index136] = ((_802_v26)[_module.__default.safeIndex(p2, new BigNumber((_802_v26).length))]).isLessThan(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of((_this).f18), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.of((_this).f18)).length)), ((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))]))).cardinality()));
          r1 = p1;
          let _805_v29;
          let _nw151 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _805_v29 = _nw151;
          let _index137 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_805_v29).length));
          (_805_v29)[_index137] = (p1).minus(p1);
          let _index138 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_805_v29).length));
          (_805_v29)[_index138] = (_this).f19;
          let _806_v30;
          let _init19 = function (_807_i2) {
            return _dafny.Set.fromElements((_this).f21);
          };
          let _nw152 = Array((new BigNumber(9)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw152.length); _i0_19++) {
            _nw152[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _806_v30 = _nw152;
          let _808_v31;
          _808_v31 = _dafny.Set.fromElements((_this).f21);
          let _index139 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_806_v30).length));
          (_806_v30)[_index139] = _808_v31;
          let _809_v32;
          _809_v32 = _dafny.Seq.UnicodeFromString("xfquxegiu");
          let _810_v33;
          _810_v33 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))],((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))]);
          let _index140 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_806_v30).length));
          let _rhs147 = _784_v8;
          let _rhs148 = _module.__default.fm16((_this).f18, new BigNumber((_module.__default.fm10((_this).f19, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(803)), ((_811_v8) => function (_812_i3) {
            return _811_v8;
          })(_784_v8))).length), globalState)).length), !((_this).f18), globalState);
          let _rhs149 = (((_this).f18) ? (_808_v31) : (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("hexqws"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(796)), ((_813_v8) => function (_814_i4) {
            return _813_v8;
          })(_784_v8)), (_this).f21, (_this).f21, _dafny.Seq.UnicodeFromString("eqel"))));
          let _rhs150 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(773)), ((_815_v8) => function (_816_i5) {
            return _815_v8;
          })(_784_v8)), _809_v32);
          let _rhs151 = (((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))]) && ((((_810_v33).contains((_this).f18)) ? ((_810_v33).get((_this).f18)) : (((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))])));
          let _lhs108 = globalState;
          let _lhs109 = _806_v30;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_806_v30).length));
          let _lhs111 = globalState;
          _784_v8 = _rhs147;
          _lhs108.f7 = _rhs148;
          _lhs109[_lhs110] = _rhs149;
          _809_v32 = _rhs150;
          _lhs111.f7 = _rhs151;
        } else {
          (globalState).f7 = !(((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))]) || ((_this).f18);
          r1 = _module.__default.fm0(globalState);
          let _817_v34;
          let _818_v35;
          let _819_v36;
          let _820_v37;
          let _out30;
          let _out31;
          let _out32;
          let _out33;
          let _outcollector10 = (_803_v27).m2((_this).f18, ((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))], ((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))], new BigNumber((((((_this).f20)[_module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f20).length))]) ? ((_this).f21) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-324)), ((_821_v8) => function (_822_i6) {
            return _821_v8;
          })(_784_v8))))).length), globalState);
          _out30 = _outcollector10[0];
          _out31 = _outcollector10[1];
          _out32 = _outcollector10[2];
          _out33 = _outcollector10[3];
          _817_v34 = _out30;
          _818_v35 = _out31;
          _819_v36 = _out32;
          _820_v37 = _out33;
          r1 = (new BigNumber(724)).minus(p2);
          let _823_v38;
          let _nw153 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _823_v38 = _nw153;
          let _index141 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_823_v38).length));
          (_823_v38)[_index141] = _820_v37;
          let _index142 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_823_v38).length));
          let _rhs152 = _module.__default.safeModuloInt(((_819_v36) ? (p1) : (new BigNumber((_783_v7).length))), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(583)), ((_824_v8) => function (_825_i7) {
            return _824_v8;
          })(_784_v8))).length));
          let _rhs153 = new BigNumber(((_this).f21).length);
          let _rhs154 = ((_this).f19).multipliedBy(p1);
          let _lhs112 = _823_v38;
          let _lhs113 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_823_v38).length));
          _lhs112[_lhs113] = _rhs152;
          _820_v37 = _rhs153;
          r1 = _rhs154;
        }
      }
      let _826_v39;
      let _nw154 = new _module.C2();
      _nw154.__ctor((_this).f18, p2);
      _826_v39 = _nw154;
      let _827_v40;
      let _nw155 = new _module.C1();
      _nw155.__ctor((_this).f18, (_this).f18, (_this).f18, p2);
      _827_v40 = _nw155;
      let _828_v41;
      _828_v41 = _dafny.Set.fromElements(_827_v40, _827_v40, _827_v40, _827_v40, _827_v40);
      let _829_v42;
      let _init20 = ((_830_p2) => function (_831_i8) {
        return _module.__default.safeDivisionInt(_831_i8, _830_p2);
      })(p2);
      let _nw156 = Array((new BigNumber(29)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw156.length); _i0_20++) {
        _nw156[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _829_v42 = _nw156;
      let _832_v43;
      _832_v43 = _dafny.Set.fromElements(_829_v42);
      let _833_v44;
      _833_v44 = _module.D8.create_DC19((_this).f20, (_this).f18, _828_v41, _832_v43, (_this).f18);
      let _834_v45;
      _834_v45 = _dafny.Map.Empty.slice().updateUnsafe(_833_v44,p2);
      _834_v45 = (_834_v45).update(_833_v44, _module.__default.fm0(globalState));
      let _835_v46;
      _835_v46 = _module.D3.create_DC7(_827_v40, _783_v7, (_this).f18, p0, (_827_v40).f18);
      let _836_v47;
      _836_v47 = _dafny.Map.Empty.slice().updateUnsafe((_835_v46).dtor_cf9,((_774_v0)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_774_v0).length))])[_module.__default.safeIndex((_this).f19, new BigNumber(((_774_v0)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_774_v0).length))]).length))]);
      let _837_v48;
      _837_v48 = _dafny.Seq.of(_module.D8.create_DC18(_836_v47));
      let _838_v49;
      _838_v49 = _module.D5.create_DC11((_827_v40).f18, _module.__default.fm0(globalState), (_this).f18, _dafny.Seq.update((_this).f21, _module.__default.safeIndex((_827_v40).f19, new BigNumber(((_this).f21).length)), new _dafny.CodePoint('c'.codePointAt(0))), (_827_v40).f19);
      let _839_v50;
      _839_v50 = _module.D8.create_DC18(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18));
      r0 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_837_v48, _837_v48), _dafny.Seq.update(_837_v48, _module.__default.safeIndex((_838_v49).dtor_cf16, new BigNumber((_837_v48).length)), _839_v50));
      r1 = (_dafny.ZERO).minus(p1);
      return [r0, r1];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
